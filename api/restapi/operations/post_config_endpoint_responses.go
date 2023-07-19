// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cybwan/fsmxlb/api/models"
)

// PostConfigEndpointNoContentCode is the HTTP code returned for type PostConfigEndpointNoContent
const PostConfigEndpointNoContentCode int = 204

/*
PostConfigEndpointNoContent OK

swagger:response postConfigEndpointNoContent
*/
type PostConfigEndpointNoContent struct {
}

// NewPostConfigEndpointNoContent creates PostConfigEndpointNoContent with default headers values
func NewPostConfigEndpointNoContent() *PostConfigEndpointNoContent {

	return &PostConfigEndpointNoContent{}
}

// WriteResponse to the client
func (o *PostConfigEndpointNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigEndpointBadRequestCode is the HTTP code returned for type PostConfigEndpointBadRequest
const PostConfigEndpointBadRequestCode int = 400

/*
PostConfigEndpointBadRequest Malformed arguments for API call

swagger:response postConfigEndpointBadRequest
*/
type PostConfigEndpointBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigEndpointBadRequest creates PostConfigEndpointBadRequest with default headers values
func NewPostConfigEndpointBadRequest() *PostConfigEndpointBadRequest {

	return &PostConfigEndpointBadRequest{}
}

// WithPayload adds the payload to the post config endpoint bad request response
func (o *PostConfigEndpointBadRequest) WithPayload(payload *models.Error) *PostConfigEndpointBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config endpoint bad request response
func (o *PostConfigEndpointBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigEndpointBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigEndpointUnauthorizedCode is the HTTP code returned for type PostConfigEndpointUnauthorized
const PostConfigEndpointUnauthorizedCode int = 401

/*
PostConfigEndpointUnauthorized Invalid authentication credentials

swagger:response postConfigEndpointUnauthorized
*/
type PostConfigEndpointUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigEndpointUnauthorized creates PostConfigEndpointUnauthorized with default headers values
func NewPostConfigEndpointUnauthorized() *PostConfigEndpointUnauthorized {

	return &PostConfigEndpointUnauthorized{}
}

// WithPayload adds the payload to the post config endpoint unauthorized response
func (o *PostConfigEndpointUnauthorized) WithPayload(payload *models.Error) *PostConfigEndpointUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config endpoint unauthorized response
func (o *PostConfigEndpointUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigEndpointUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigEndpointForbiddenCode is the HTTP code returned for type PostConfigEndpointForbidden
const PostConfigEndpointForbiddenCode int = 403

/*
PostConfigEndpointForbidden Capacity insufficient

swagger:response postConfigEndpointForbidden
*/
type PostConfigEndpointForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigEndpointForbidden creates PostConfigEndpointForbidden with default headers values
func NewPostConfigEndpointForbidden() *PostConfigEndpointForbidden {

	return &PostConfigEndpointForbidden{}
}

// WithPayload adds the payload to the post config endpoint forbidden response
func (o *PostConfigEndpointForbidden) WithPayload(payload *models.Error) *PostConfigEndpointForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config endpoint forbidden response
func (o *PostConfigEndpointForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigEndpointForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigEndpointNotFoundCode is the HTTP code returned for type PostConfigEndpointNotFound
const PostConfigEndpointNotFoundCode int = 404

/*
PostConfigEndpointNotFound Resource not found

swagger:response postConfigEndpointNotFound
*/
type PostConfigEndpointNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigEndpointNotFound creates PostConfigEndpointNotFound with default headers values
func NewPostConfigEndpointNotFound() *PostConfigEndpointNotFound {

	return &PostConfigEndpointNotFound{}
}

// WithPayload adds the payload to the post config endpoint not found response
func (o *PostConfigEndpointNotFound) WithPayload(payload *models.Error) *PostConfigEndpointNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config endpoint not found response
func (o *PostConfigEndpointNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigEndpointNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigEndpointConflictCode is the HTTP code returned for type PostConfigEndpointConflict
const PostConfigEndpointConflictCode int = 409

/*
PostConfigEndpointConflict Resource Conflict.

swagger:response postConfigEndpointConflict
*/
type PostConfigEndpointConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigEndpointConflict creates PostConfigEndpointConflict with default headers values
func NewPostConfigEndpointConflict() *PostConfigEndpointConflict {

	return &PostConfigEndpointConflict{}
}

// WithPayload adds the payload to the post config endpoint conflict response
func (o *PostConfigEndpointConflict) WithPayload(payload *models.Error) *PostConfigEndpointConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config endpoint conflict response
func (o *PostConfigEndpointConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigEndpointConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigEndpointInternalServerErrorCode is the HTTP code returned for type PostConfigEndpointInternalServerError
const PostConfigEndpointInternalServerErrorCode int = 500

/*
PostConfigEndpointInternalServerError Internal service error

swagger:response postConfigEndpointInternalServerError
*/
type PostConfigEndpointInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigEndpointInternalServerError creates PostConfigEndpointInternalServerError with default headers values
func NewPostConfigEndpointInternalServerError() *PostConfigEndpointInternalServerError {

	return &PostConfigEndpointInternalServerError{}
}

// WithPayload adds the payload to the post config endpoint internal server error response
func (o *PostConfigEndpointInternalServerError) WithPayload(payload *models.Error) *PostConfigEndpointInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config endpoint internal server error response
func (o *PostConfigEndpointInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigEndpointInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigEndpointServiceUnavailableCode is the HTTP code returned for type PostConfigEndpointServiceUnavailable
const PostConfigEndpointServiceUnavailableCode int = 503

/*
PostConfigEndpointServiceUnavailable Maintanence mode

swagger:response postConfigEndpointServiceUnavailable
*/
type PostConfigEndpointServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigEndpointServiceUnavailable creates PostConfigEndpointServiceUnavailable with default headers values
func NewPostConfigEndpointServiceUnavailable() *PostConfigEndpointServiceUnavailable {

	return &PostConfigEndpointServiceUnavailable{}
}

// WithPayload adds the payload to the post config endpoint service unavailable response
func (o *PostConfigEndpointServiceUnavailable) WithPayload(payload *models.Error) *PostConfigEndpointServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config endpoint service unavailable response
func (o *PostConfigEndpointServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigEndpointServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
