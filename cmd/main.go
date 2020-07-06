package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/filanov/bm-inventory/internal/versions"

	"github.com/filanov/bm-inventory/internal/bminventory"
	"github.com/filanov/bm-inventory/internal/cluster"
	"github.com/filanov/bm-inventory/internal/common"
	"github.com/filanov/bm-inventory/internal/metrics"

	"github.com/filanov/bm-inventory/internal/events"
	"github.com/filanov/bm-inventory/internal/hardware"
	"github.com/filanov/bm-inventory/internal/host"
	"github.com/filanov/bm-inventory/models"
	"github.com/filanov/bm-inventory/pkg/app"
	"github.com/filanov/bm-inventory/pkg/job"
	"github.com/filanov/bm-inventory/pkg/requestid"
	awsS3Client "github.com/filanov/bm-inventory/pkg/s3Client"
	"github.com/filanov/bm-inventory/pkg/s3wrapper"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	"github.com/filanov/bm-inventory/pkg/thread"
	"github.com/filanov/bm-inventory/restapi"
)

func init() {
	strfmt.MarshalFormat = strfmt.ISO8601LocalTime
}

var Options struct {
	BMConfig                    bminventory.Config
	DBHost                      string `envconfig:"DB_HOST" default:"mariadb"`
	DBPort                      string `envconfig:"DB_PORT" default:"3306"`
	HWValidatorConfig           hardware.ValidatorCfg
	JobConfig                   job.Config
	InstructionConfig           host.InstructionConfig
	ClusterStateMonitorInterval time.Duration `envconfig:"CLUSTER_MONITOR_INTERVAL" default:"10s"`
	S3Config                    s3wrapper.Config
	HostStateMonitorInterval    time.Duration `envconfig:"HOST_MONITOR_INTERVAL" default:"30s"`
	Versions                    versions.Versions
	UseK8s                      bool   `envconfig:"USE_K8S" default:"true"` // TODO remove when jobs running deprecated
	Target                      string `envconfig:"TARGET" default:""`
}

func main() {
	log := logrus.New()
	log.SetReportCaller(true)

	err := envconfig.Process("myapp", &Options)
	if err != nil {
		log.Fatal(err.Error())
	}

	port := flag.String("port", "8090", "define port that the service will listen to")
	flag.Parse()

	log.Println("Starting bm service")

	db, err := gorm.Open("mysql",
		fmt.Sprintf("admin:admin@tcp(%s:%s)/installer?charset=utf8&parseTime=True&loc=Local",
			Options.DBHost, Options.DBPort))

	if err != nil {
		log.Fatal("Fail to connect to DB, ", err)
	}
	defer db.Close()
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(0)
	db.DB().SetConnMaxLifetime(0)

	if err = db.AutoMigrate(&models.Host{}, &common.Cluster{}, &events.Event{}).Error; err != nil {
		log.Fatal("failed to auto migrate, ", err)
	}

	versionHandler := versions.NewHandler(Options.Versions)
	eventsHandler := events.New(db, log.WithField("pkg", "events"))
	hwValidator := hardware.NewValidator(log.WithField("pkg", "validators"), Options.HWValidatorConfig)
	instructionApi := host.NewInstructionManager(log, db, hwValidator, Options.InstructionConfig)
	hostApi := host.NewManager(log.WithField("pkg", "host-state"), db, hwValidator, instructionApi)
	clusterApi := cluster.NewManager(log.WithField("pkg", "cluster-state"), db, eventsHandler)

	clusterStateMonitor := thread.New(
		log.WithField("pkg", "cluster-monitor"), "Cluster State Monitor", Options.ClusterStateMonitorInterval, clusterApi.ClusterMonitoring)
	clusterStateMonitor.Start()
	defer clusterStateMonitor.Stop()

	hostStateMonitor := thread.New(
		log.WithField("pkg", "host-monitor"), "Host State Monitor", Options.HostStateMonitorInterval, hostApi.HostMonitoring)
	hostStateMonitor.Start()
	defer hostStateMonitor.Stop()

	s3Client, err := awsS3Client.NewS3Client(Options.BMConfig.S3EndpointURL, Options.BMConfig.AwsAccessKeyID, Options.BMConfig.AwsSecretAccessKey, log)
	if err != nil {
		log.Fatal("Failed to setup S3 client", err)
	}

	var generator job.ISOInstallConfigGenerator

	if Options.Target != "disconnected" {
		var kclient client.Client
		if Options.UseK8s {
			createS3Bucket(&Options.S3Config)

			scheme := runtime.NewScheme()
			if err = clientgoscheme.AddToScheme(scheme); err != nil {
				log.Fatal("Failed to add K8S scheme", err)
			}

			kclient, err = client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme})
			if err != nil && Options.UseK8s {
				log.Fatal("failed to create client:", err)
			}

		} else {
			log.Println("running drone test, skipping S3")
			kclient = nil
		}

		generator = job.New(log.WithField("pkg", "k8s-job-wrapper"), kclient, Options.JobConfig)
	} else {
		// in disconnected mode, setup s3 and use localjob implementation
		createS3Bucket(&Options.S3Config)
		generator = job.NewLocalJob(log.WithField("pkg", "local-job-wrapper"), Options.JobConfig)
	}

	bm := bminventory.NewBareMetalInventory(db, log.WithField("pkg", "Inventory"), hostApi, clusterApi, Options.BMConfig, generator, eventsHandler, s3Client)

	events := events.NewApi(eventsHandler, logrus.WithField("pkg", "eventsApi"))

	h, err := restapi.Handler(restapi.Config{
		InstallerAPI:    bm,
		EventsAPI:       events,
		Logger:          log.Printf,
		VersionsAPI:     versionHandler,
		InnerMiddleware: metrics.WithMatchedRoute(log.WithField("pkg", "matched-h")),
	})
	h = app.WithMetricsResponderMiddleware(h)
	h = app.WithHealthMiddleware(h)

	h = requestid.Middleware(h)
	if err != nil {
		log.Fatal("Failed to init rest handler,", err)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", swag.StringValue(port)), h))
}

func createS3Bucket(s3config *s3wrapper.Config) {
	if err := s3wrapper.CreateBucket(&Options.S3Config); err != nil {
		log.Fatal(err)
	}
}
