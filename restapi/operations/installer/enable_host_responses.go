// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/filanov/bm-inventory/models"
)

// EnableHostNoContentCode is the HTTP code returned for type EnableHostNoContent
const EnableHostNoContentCode int = 204

/*EnableHostNoContent Success.

swagger:response enableHostNoContent
*/
type EnableHostNoContent struct {
}

// NewEnableHostNoContent creates EnableHostNoContent with default headers values
func NewEnableHostNoContent() *EnableHostNoContent {

	return &EnableHostNoContent{}
}

// WriteResponse to the client
func (o *EnableHostNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// EnableHostNotFoundCode is the HTTP code returned for type EnableHostNotFound
const EnableHostNotFoundCode int = 404

/*EnableHostNotFound Error.

swagger:response enableHostNotFound
*/
type EnableHostNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEnableHostNotFound creates EnableHostNotFound with default headers values
func NewEnableHostNotFound() *EnableHostNotFound {

	return &EnableHostNotFound{}
}

// WithPayload adds the payload to the enable host not found response
func (o *EnableHostNotFound) WithPayload(payload *models.Error) *EnableHostNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable host not found response
func (o *EnableHostNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableHostNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableHostConflictCode is the HTTP code returned for type EnableHostConflict
const EnableHostConflictCode int = 409

/*EnableHostConflict Error.

swagger:response enableHostConflict
*/
type EnableHostConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEnableHostConflict creates EnableHostConflict with default headers values
func NewEnableHostConflict() *EnableHostConflict {

	return &EnableHostConflict{}
}

// WithPayload adds the payload to the enable host conflict response
func (o *EnableHostConflict) WithPayload(payload *models.Error) *EnableHostConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable host conflict response
func (o *EnableHostConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableHostConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableHostInternalServerErrorCode is the HTTP code returned for type EnableHostInternalServerError
const EnableHostInternalServerErrorCode int = 500

/*EnableHostInternalServerError Error.

swagger:response enableHostInternalServerError
*/
type EnableHostInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEnableHostInternalServerError creates EnableHostInternalServerError with default headers values
func NewEnableHostInternalServerError() *EnableHostInternalServerError {

	return &EnableHostInternalServerError{}
}

// WithPayload adds the payload to the enable host internal server error response
func (o *EnableHostInternalServerError) WithPayload(payload *models.Error) *EnableHostInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable host internal server error response
func (o *EnableHostInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableHostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}