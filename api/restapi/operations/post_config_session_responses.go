// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/flomesh-io/fnlb/api/models"
)

// PostConfigSessionNoContentCode is the HTTP code returned for type PostConfigSessionNoContent
const PostConfigSessionNoContentCode int = 204

/*
PostConfigSessionNoContent OK

swagger:response postConfigSessionNoContent
*/
type PostConfigSessionNoContent struct {
}

// NewPostConfigSessionNoContent creates PostConfigSessionNoContent with default headers values
func NewPostConfigSessionNoContent() *PostConfigSessionNoContent {

	return &PostConfigSessionNoContent{}
}

// WriteResponse to the client
func (o *PostConfigSessionNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigSessionBadRequestCode is the HTTP code returned for type PostConfigSessionBadRequest
const PostConfigSessionBadRequestCode int = 400

/*
PostConfigSessionBadRequest Malformed arguments for API call

swagger:response postConfigSessionBadRequest
*/
type PostConfigSessionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionBadRequest creates PostConfigSessionBadRequest with default headers values
func NewPostConfigSessionBadRequest() *PostConfigSessionBadRequest {

	return &PostConfigSessionBadRequest{}
}

// WithPayload adds the payload to the post config session bad request response
func (o *PostConfigSessionBadRequest) WithPayload(payload *models.Error) *PostConfigSessionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config session bad request response
func (o *PostConfigSessionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionUnauthorizedCode is the HTTP code returned for type PostConfigSessionUnauthorized
const PostConfigSessionUnauthorizedCode int = 401

/*
PostConfigSessionUnauthorized Invalid authentication credentials

swagger:response postConfigSessionUnauthorized
*/
type PostConfigSessionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionUnauthorized creates PostConfigSessionUnauthorized with default headers values
func NewPostConfigSessionUnauthorized() *PostConfigSessionUnauthorized {

	return &PostConfigSessionUnauthorized{}
}

// WithPayload adds the payload to the post config session unauthorized response
func (o *PostConfigSessionUnauthorized) WithPayload(payload *models.Error) *PostConfigSessionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config session unauthorized response
func (o *PostConfigSessionUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionForbiddenCode is the HTTP code returned for type PostConfigSessionForbidden
const PostConfigSessionForbiddenCode int = 403

/*
PostConfigSessionForbidden Capacity insufficient

swagger:response postConfigSessionForbidden
*/
type PostConfigSessionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionForbidden creates PostConfigSessionForbidden with default headers values
func NewPostConfigSessionForbidden() *PostConfigSessionForbidden {

	return &PostConfigSessionForbidden{}
}

// WithPayload adds the payload to the post config session forbidden response
func (o *PostConfigSessionForbidden) WithPayload(payload *models.Error) *PostConfigSessionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config session forbidden response
func (o *PostConfigSessionForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionNotFoundCode is the HTTP code returned for type PostConfigSessionNotFound
const PostConfigSessionNotFoundCode int = 404

/*
PostConfigSessionNotFound Resource not found

swagger:response postConfigSessionNotFound
*/
type PostConfigSessionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionNotFound creates PostConfigSessionNotFound with default headers values
func NewPostConfigSessionNotFound() *PostConfigSessionNotFound {

	return &PostConfigSessionNotFound{}
}

// WithPayload adds the payload to the post config session not found response
func (o *PostConfigSessionNotFound) WithPayload(payload *models.Error) *PostConfigSessionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config session not found response
func (o *PostConfigSessionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionConflictCode is the HTTP code returned for type PostConfigSessionConflict
const PostConfigSessionConflictCode int = 409

/*
PostConfigSessionConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response postConfigSessionConflict
*/
type PostConfigSessionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionConflict creates PostConfigSessionConflict with default headers values
func NewPostConfigSessionConflict() *PostConfigSessionConflict {

	return &PostConfigSessionConflict{}
}

// WithPayload adds the payload to the post config session conflict response
func (o *PostConfigSessionConflict) WithPayload(payload *models.Error) *PostConfigSessionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config session conflict response
func (o *PostConfigSessionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionInternalServerErrorCode is the HTTP code returned for type PostConfigSessionInternalServerError
const PostConfigSessionInternalServerErrorCode int = 500

/*
PostConfigSessionInternalServerError Internal service error

swagger:response postConfigSessionInternalServerError
*/
type PostConfigSessionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionInternalServerError creates PostConfigSessionInternalServerError with default headers values
func NewPostConfigSessionInternalServerError() *PostConfigSessionInternalServerError {

	return &PostConfigSessionInternalServerError{}
}

// WithPayload adds the payload to the post config session internal server error response
func (o *PostConfigSessionInternalServerError) WithPayload(payload *models.Error) *PostConfigSessionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config session internal server error response
func (o *PostConfigSessionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionServiceUnavailableCode is the HTTP code returned for type PostConfigSessionServiceUnavailable
const PostConfigSessionServiceUnavailableCode int = 503

/*
PostConfigSessionServiceUnavailable Maintanence mode

swagger:response postConfigSessionServiceUnavailable
*/
type PostConfigSessionServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionServiceUnavailable creates PostConfigSessionServiceUnavailable with default headers values
func NewPostConfigSessionServiceUnavailable() *PostConfigSessionServiceUnavailable {

	return &PostConfigSessionServiceUnavailable{}
}

// WithPayload adds the payload to the post config session service unavailable response
func (o *PostConfigSessionServiceUnavailable) WithPayload(payload *models.Error) *PostConfigSessionServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config session service unavailable response
func (o *PostConfigSessionServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
