// Code generated by MockGen. DO NOT EDIT.
// Source: cluster.go

// Package cluster is a generated GoMock package.
package cluster

import (
	context "context"
	reflect "reflect"

	models "github.com/filanov/bm-inventory/models"
	strfmt "github.com/go-openapi/strfmt"
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
)

// MockStateAPI is a mock of StateAPI interface.
type MockStateAPI struct {
	ctrl     *gomock.Controller
	recorder *MockStateAPIMockRecorder
}

// MockStateAPIMockRecorder is the mock recorder for MockStateAPI.
type MockStateAPIMockRecorder struct {
	mock *MockStateAPI
}

// NewMockStateAPI creates a new mock instance.
func NewMockStateAPI(ctrl *gomock.Controller) *MockStateAPI {
	mock := &MockStateAPI{ctrl: ctrl}
	mock.recorder = &MockStateAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateAPI) EXPECT() *MockStateAPIMockRecorder {
	return m.recorder
}

// RefreshStatus mocks base method.
func (m *MockStateAPI) RefreshStatus(ctx context.Context, c *models.Cluster, db *gorm.DB) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshStatus", ctx, c, db)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshStatus indicates an expected call of RefreshStatus.
func (mr *MockStateAPIMockRecorder) RefreshStatus(ctx, c, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshStatus", reflect.TypeOf((*MockStateAPI)(nil).RefreshStatus), ctx, c, db)
}

// MockRegistrationAPI is a mock of RegistrationAPI interface.
type MockRegistrationAPI struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrationAPIMockRecorder
}

// MockRegistrationAPIMockRecorder is the mock recorder for MockRegistrationAPI.
type MockRegistrationAPIMockRecorder struct {
	mock *MockRegistrationAPI
}

// NewMockRegistrationAPI creates a new mock instance.
func NewMockRegistrationAPI(ctrl *gomock.Controller) *MockRegistrationAPI {
	mock := &MockRegistrationAPI{ctrl: ctrl}
	mock.recorder = &MockRegistrationAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistrationAPI) EXPECT() *MockRegistrationAPIMockRecorder {
	return m.recorder
}

// RegisterCluster mocks base method.
func (m *MockRegistrationAPI) RegisterCluster(ctx context.Context, c *models.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCluster", ctx, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterCluster indicates an expected call of RegisterCluster.
func (mr *MockRegistrationAPIMockRecorder) RegisterCluster(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCluster", reflect.TypeOf((*MockRegistrationAPI)(nil).RegisterCluster), ctx, c)
}

// DeregisterCluster mocks base method.
func (m *MockRegistrationAPI) DeregisterCluster(ctx context.Context, c *models.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterCluster", ctx, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeregisterCluster indicates an expected call of DeregisterCluster.
func (mr *MockRegistrationAPIMockRecorder) DeregisterCluster(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterCluster", reflect.TypeOf((*MockRegistrationAPI)(nil).DeregisterCluster), ctx, c)
}

// MockInstallationAPI is a mock of InstallationAPI interface.
type MockInstallationAPI struct {
	ctrl     *gomock.Controller
	recorder *MockInstallationAPIMockRecorder
}

// MockInstallationAPIMockRecorder is the mock recorder for MockInstallationAPI.
type MockInstallationAPIMockRecorder struct {
	mock *MockInstallationAPI
}

// NewMockInstallationAPI creates a new mock instance.
func NewMockInstallationAPI(ctrl *gomock.Controller) *MockInstallationAPI {
	mock := &MockInstallationAPI{ctrl: ctrl}
	mock.recorder = &MockInstallationAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstallationAPI) EXPECT() *MockInstallationAPIMockRecorder {
	return m.recorder
}

// Install mocks base method.
func (m *MockInstallationAPI) Install(ctx context.Context, c *models.Cluster, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", ctx, c, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockInstallationAPIMockRecorder) Install(ctx, c, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockInstallationAPI)(nil).Install), ctx, c, db)
}

// GetMasterNodesIds mocks base method.
func (m *MockInstallationAPI) GetMasterNodesIds(ctx context.Context, c *models.Cluster, db *gorm.DB) ([]*strfmt.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMasterNodesIds", ctx, c, db)
	ret0, _ := ret[0].([]*strfmt.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMasterNodesIds indicates an expected call of GetMasterNodesIds.
func (mr *MockInstallationAPIMockRecorder) GetMasterNodesIds(ctx, c, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMasterNodesIds", reflect.TypeOf((*MockInstallationAPI)(nil).GetMasterNodesIds), ctx, c, db)
}

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// RefreshStatus mocks base method.
func (m *MockAPI) RefreshStatus(ctx context.Context, c *models.Cluster, db *gorm.DB) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshStatus", ctx, c, db)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshStatus indicates an expected call of RefreshStatus.
func (mr *MockAPIMockRecorder) RefreshStatus(ctx, c, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshStatus", reflect.TypeOf((*MockAPI)(nil).RefreshStatus), ctx, c, db)
}

// RegisterCluster mocks base method.
func (m *MockAPI) RegisterCluster(ctx context.Context, c *models.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCluster", ctx, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterCluster indicates an expected call of RegisterCluster.
func (mr *MockAPIMockRecorder) RegisterCluster(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCluster", reflect.TypeOf((*MockAPI)(nil).RegisterCluster), ctx, c)
}

// DeregisterCluster mocks base method.
func (m *MockAPI) DeregisterCluster(ctx context.Context, c *models.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterCluster", ctx, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeregisterCluster indicates an expected call of DeregisterCluster.
func (mr *MockAPIMockRecorder) DeregisterCluster(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterCluster", reflect.TypeOf((*MockAPI)(nil).DeregisterCluster), ctx, c)
}

// Install mocks base method.
func (m *MockAPI) Install(ctx context.Context, c *models.Cluster, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", ctx, c, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockAPIMockRecorder) Install(ctx, c, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockAPI)(nil).Install), ctx, c, db)
}

// GetMasterNodesIds mocks base method.
func (m *MockAPI) GetMasterNodesIds(ctx context.Context, c *models.Cluster, db *gorm.DB) ([]*strfmt.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMasterNodesIds", ctx, c, db)
	ret0, _ := ret[0].([]*strfmt.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMasterNodesIds indicates an expected call of GetMasterNodesIds.
func (mr *MockAPIMockRecorder) GetMasterNodesIds(ctx, c, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMasterNodesIds", reflect.TypeOf((*MockAPI)(nil).GetMasterNodesIds), ctx, c, db)
}

// ClusterMonitoring mocks base method.
func (m *MockAPI) ClusterMonitoring() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterMonitoring")
}

// ClusterMonitoring indicates an expected call of ClusterMonitoring.
func (mr *MockAPIMockRecorder) ClusterMonitoring() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterMonitoring", reflect.TypeOf((*MockAPI)(nil).ClusterMonitoring))
}
