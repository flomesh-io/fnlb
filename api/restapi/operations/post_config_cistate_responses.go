// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cybwan/fsmxlb/api/models"
)

// PostConfigCistateNoContentCode is the HTTP code returned for type PostConfigCistateNoContent
const PostConfigCistateNoContentCode int = 204

/*
PostConfigCistateNoContent OK

swagger:response postConfigCistateNoContent
*/
type PostConfigCistateNoContent struct {
}

// NewPostConfigCistateNoContent creates PostConfigCistateNoContent with default headers values
func NewPostConfigCistateNoContent() *PostConfigCistateNoContent {

	return &PostConfigCistateNoContent{}
}

// WriteResponse to the client
func (o *PostConfigCistateNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigCistateBadRequestCode is the HTTP code returned for type PostConfigCistateBadRequest
const PostConfigCistateBadRequestCode int = 400

/*
PostConfigCistateBadRequest Malformed arguments for API call

swagger:response postConfigCistateBadRequest
*/
type PostConfigCistateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigCistateBadRequest creates PostConfigCistateBadRequest with default headers values
func NewPostConfigCistateBadRequest() *PostConfigCistateBadRequest {

	return &PostConfigCistateBadRequest{}
}

// WithPayload adds the payload to the post config cistate bad request response
func (o *PostConfigCistateBadRequest) WithPayload(payload *models.Error) *PostConfigCistateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config cistate bad request response
func (o *PostConfigCistateBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigCistateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigCistateUnauthorizedCode is the HTTP code returned for type PostConfigCistateUnauthorized
const PostConfigCistateUnauthorizedCode int = 401

/*
PostConfigCistateUnauthorized Invalid authentication credentials

swagger:response postConfigCistateUnauthorized
*/
type PostConfigCistateUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigCistateUnauthorized creates PostConfigCistateUnauthorized with default headers values
func NewPostConfigCistateUnauthorized() *PostConfigCistateUnauthorized {

	return &PostConfigCistateUnauthorized{}
}

// WithPayload adds the payload to the post config cistate unauthorized response
func (o *PostConfigCistateUnauthorized) WithPayload(payload *models.Error) *PostConfigCistateUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config cistate unauthorized response
func (o *PostConfigCistateUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigCistateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigCistateForbiddenCode is the HTTP code returned for type PostConfigCistateForbidden
const PostConfigCistateForbiddenCode int = 403

/*
PostConfigCistateForbidden Capacity insufficient

swagger:response postConfigCistateForbidden
*/
type PostConfigCistateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigCistateForbidden creates PostConfigCistateForbidden with default headers values
func NewPostConfigCistateForbidden() *PostConfigCistateForbidden {

	return &PostConfigCistateForbidden{}
}

// WithPayload adds the payload to the post config cistate forbidden response
func (o *PostConfigCistateForbidden) WithPayload(payload *models.Error) *PostConfigCistateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config cistate forbidden response
func (o *PostConfigCistateForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigCistateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigCistateNotFoundCode is the HTTP code returned for type PostConfigCistateNotFound
const PostConfigCistateNotFoundCode int = 404

/*
PostConfigCistateNotFound Resource not found

swagger:response postConfigCistateNotFound
*/
type PostConfigCistateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigCistateNotFound creates PostConfigCistateNotFound with default headers values
func NewPostConfigCistateNotFound() *PostConfigCistateNotFound {

	return &PostConfigCistateNotFound{}
}

// WithPayload adds the payload to the post config cistate not found response
func (o *PostConfigCistateNotFound) WithPayload(payload *models.Error) *PostConfigCistateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config cistate not found response
func (o *PostConfigCistateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigCistateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigCistateConflictCode is the HTTP code returned for type PostConfigCistateConflict
const PostConfigCistateConflictCode int = 409

/*
PostConfigCistateConflict Resource Conflict.

swagger:response postConfigCistateConflict
*/
type PostConfigCistateConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigCistateConflict creates PostConfigCistateConflict with default headers values
func NewPostConfigCistateConflict() *PostConfigCistateConflict {

	return &PostConfigCistateConflict{}
}

// WithPayload adds the payload to the post config cistate conflict response
func (o *PostConfigCistateConflict) WithPayload(payload *models.Error) *PostConfigCistateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config cistate conflict response
func (o *PostConfigCistateConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigCistateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigCistateInternalServerErrorCode is the HTTP code returned for type PostConfigCistateInternalServerError
const PostConfigCistateInternalServerErrorCode int = 500

/*
PostConfigCistateInternalServerError Internal service error

swagger:response postConfigCistateInternalServerError
*/
type PostConfigCistateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigCistateInternalServerError creates PostConfigCistateInternalServerError with default headers values
func NewPostConfigCistateInternalServerError() *PostConfigCistateInternalServerError {

	return &PostConfigCistateInternalServerError{}
}

// WithPayload adds the payload to the post config cistate internal server error response
func (o *PostConfigCistateInternalServerError) WithPayload(payload *models.Error) *PostConfigCistateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config cistate internal server error response
func (o *PostConfigCistateInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigCistateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigCistateServiceUnavailableCode is the HTTP code returned for type PostConfigCistateServiceUnavailable
const PostConfigCistateServiceUnavailableCode int = 503

/*
PostConfigCistateServiceUnavailable Maintanence mode

swagger:response postConfigCistateServiceUnavailable
*/
type PostConfigCistateServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigCistateServiceUnavailable creates PostConfigCistateServiceUnavailable with default headers values
func NewPostConfigCistateServiceUnavailable() *PostConfigCistateServiceUnavailable {

	return &PostConfigCistateServiceUnavailable{}
}

// WithPayload adds the payload to the post config cistate service unavailable response
func (o *PostConfigCistateServiceUnavailable) WithPayload(payload *models.Error) *PostConfigCistateServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config cistate service unavailable response
func (o *PostConfigCistateServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigCistateServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
