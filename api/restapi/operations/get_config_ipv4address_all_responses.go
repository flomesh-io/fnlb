// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/flomesh-io/fnlb/api/models"
)

// GetConfigIpv4addressAllOKCode is the HTTP code returned for type GetConfigIpv4addressAllOK
const GetConfigIpv4addressAllOKCode int = 200

/*
GetConfigIpv4addressAllOK OK

swagger:response getConfigIpv4addressAllOK
*/
type GetConfigIpv4addressAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigIpv4addressAllOKBody `json:"body,omitempty"`
}

// NewGetConfigIpv4addressAllOK creates GetConfigIpv4addressAllOK with default headers values
func NewGetConfigIpv4addressAllOK() *GetConfigIpv4addressAllOK {

	return &GetConfigIpv4addressAllOK{}
}

// WithPayload adds the payload to the get config ipv4address all o k response
func (o *GetConfigIpv4addressAllOK) WithPayload(payload *GetConfigIpv4addressAllOKBody) *GetConfigIpv4addressAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config ipv4address all o k response
func (o *GetConfigIpv4addressAllOK) SetPayload(payload *GetConfigIpv4addressAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigIpv4addressAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigIpv4addressAllUnauthorizedCode is the HTTP code returned for type GetConfigIpv4addressAllUnauthorized
const GetConfigIpv4addressAllUnauthorizedCode int = 401

/*
GetConfigIpv4addressAllUnauthorized Invalid authentication credentials

swagger:response getConfigIpv4addressAllUnauthorized
*/
type GetConfigIpv4addressAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigIpv4addressAllUnauthorized creates GetConfigIpv4addressAllUnauthorized with default headers values
func NewGetConfigIpv4addressAllUnauthorized() *GetConfigIpv4addressAllUnauthorized {

	return &GetConfigIpv4addressAllUnauthorized{}
}

// WithPayload adds the payload to the get config ipv4address all unauthorized response
func (o *GetConfigIpv4addressAllUnauthorized) WithPayload(payload *models.Error) *GetConfigIpv4addressAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config ipv4address all unauthorized response
func (o *GetConfigIpv4addressAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigIpv4addressAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigIpv4addressAllInternalServerErrorCode is the HTTP code returned for type GetConfigIpv4addressAllInternalServerError
const GetConfigIpv4addressAllInternalServerErrorCode int = 500

/*
GetConfigIpv4addressAllInternalServerError Internal service error

swagger:response getConfigIpv4addressAllInternalServerError
*/
type GetConfigIpv4addressAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigIpv4addressAllInternalServerError creates GetConfigIpv4addressAllInternalServerError with default headers values
func NewGetConfigIpv4addressAllInternalServerError() *GetConfigIpv4addressAllInternalServerError {

	return &GetConfigIpv4addressAllInternalServerError{}
}

// WithPayload adds the payload to the get config ipv4address all internal server error response
func (o *GetConfigIpv4addressAllInternalServerError) WithPayload(payload *models.Error) *GetConfigIpv4addressAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config ipv4address all internal server error response
func (o *GetConfigIpv4addressAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigIpv4addressAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigIpv4addressAllServiceUnavailableCode is the HTTP code returned for type GetConfigIpv4addressAllServiceUnavailable
const GetConfigIpv4addressAllServiceUnavailableCode int = 503

/*
GetConfigIpv4addressAllServiceUnavailable Maintanence mode

swagger:response getConfigIpv4addressAllServiceUnavailable
*/
type GetConfigIpv4addressAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigIpv4addressAllServiceUnavailable creates GetConfigIpv4addressAllServiceUnavailable with default headers values
func NewGetConfigIpv4addressAllServiceUnavailable() *GetConfigIpv4addressAllServiceUnavailable {

	return &GetConfigIpv4addressAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config ipv4address all service unavailable response
func (o *GetConfigIpv4addressAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigIpv4addressAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config ipv4address all service unavailable response
func (o *GetConfigIpv4addressAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigIpv4addressAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
