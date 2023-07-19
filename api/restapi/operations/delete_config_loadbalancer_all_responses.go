// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cybwan/fsmxlb/api/models"
)

// DeleteConfigLoadbalancerAllNoContentCode is the HTTP code returned for type DeleteConfigLoadbalancerAllNoContent
const DeleteConfigLoadbalancerAllNoContentCode int = 204

/*
DeleteConfigLoadbalancerAllNoContent OK

swagger:response deleteConfigLoadbalancerAllNoContent
*/
type DeleteConfigLoadbalancerAllNoContent struct {
}

// NewDeleteConfigLoadbalancerAllNoContent creates DeleteConfigLoadbalancerAllNoContent with default headers values
func NewDeleteConfigLoadbalancerAllNoContent() *DeleteConfigLoadbalancerAllNoContent {

	return &DeleteConfigLoadbalancerAllNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerAllNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigLoadbalancerAllBadRequestCode is the HTTP code returned for type DeleteConfigLoadbalancerAllBadRequest
const DeleteConfigLoadbalancerAllBadRequestCode int = 400

/*
DeleteConfigLoadbalancerAllBadRequest Malformed arguments for API call

swagger:response deleteConfigLoadbalancerAllBadRequest
*/
type DeleteConfigLoadbalancerAllBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerAllBadRequest creates DeleteConfigLoadbalancerAllBadRequest with default headers values
func NewDeleteConfigLoadbalancerAllBadRequest() *DeleteConfigLoadbalancerAllBadRequest {

	return &DeleteConfigLoadbalancerAllBadRequest{}
}

// WithPayload adds the payload to the delete config loadbalancer all bad request response
func (o *DeleteConfigLoadbalancerAllBadRequest) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerAllBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer all bad request response
func (o *DeleteConfigLoadbalancerAllBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerAllBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerAllUnauthorizedCode is the HTTP code returned for type DeleteConfigLoadbalancerAllUnauthorized
const DeleteConfigLoadbalancerAllUnauthorizedCode int = 401

/*
DeleteConfigLoadbalancerAllUnauthorized Invalid authentication credentials

swagger:response deleteConfigLoadbalancerAllUnauthorized
*/
type DeleteConfigLoadbalancerAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerAllUnauthorized creates DeleteConfigLoadbalancerAllUnauthorized with default headers values
func NewDeleteConfigLoadbalancerAllUnauthorized() *DeleteConfigLoadbalancerAllUnauthorized {

	return &DeleteConfigLoadbalancerAllUnauthorized{}
}

// WithPayload adds the payload to the delete config loadbalancer all unauthorized response
func (o *DeleteConfigLoadbalancerAllUnauthorized) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer all unauthorized response
func (o *DeleteConfigLoadbalancerAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerAllForbiddenCode is the HTTP code returned for type DeleteConfigLoadbalancerAllForbidden
const DeleteConfigLoadbalancerAllForbiddenCode int = 403

/*
DeleteConfigLoadbalancerAllForbidden Capacity insufficient

swagger:response deleteConfigLoadbalancerAllForbidden
*/
type DeleteConfigLoadbalancerAllForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerAllForbidden creates DeleteConfigLoadbalancerAllForbidden with default headers values
func NewDeleteConfigLoadbalancerAllForbidden() *DeleteConfigLoadbalancerAllForbidden {

	return &DeleteConfigLoadbalancerAllForbidden{}
}

// WithPayload adds the payload to the delete config loadbalancer all forbidden response
func (o *DeleteConfigLoadbalancerAllForbidden) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerAllForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer all forbidden response
func (o *DeleteConfigLoadbalancerAllForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerAllForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerAllNotFoundCode is the HTTP code returned for type DeleteConfigLoadbalancerAllNotFound
const DeleteConfigLoadbalancerAllNotFoundCode int = 404

/*
DeleteConfigLoadbalancerAllNotFound Resource not found

swagger:response deleteConfigLoadbalancerAllNotFound
*/
type DeleteConfigLoadbalancerAllNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerAllNotFound creates DeleteConfigLoadbalancerAllNotFound with default headers values
func NewDeleteConfigLoadbalancerAllNotFound() *DeleteConfigLoadbalancerAllNotFound {

	return &DeleteConfigLoadbalancerAllNotFound{}
}

// WithPayload adds the payload to the delete config loadbalancer all not found response
func (o *DeleteConfigLoadbalancerAllNotFound) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerAllNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer all not found response
func (o *DeleteConfigLoadbalancerAllNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerAllNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerAllConflictCode is the HTTP code returned for type DeleteConfigLoadbalancerAllConflict
const DeleteConfigLoadbalancerAllConflictCode int = 409

/*
DeleteConfigLoadbalancerAllConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response deleteConfigLoadbalancerAllConflict
*/
type DeleteConfigLoadbalancerAllConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerAllConflict creates DeleteConfigLoadbalancerAllConflict with default headers values
func NewDeleteConfigLoadbalancerAllConflict() *DeleteConfigLoadbalancerAllConflict {

	return &DeleteConfigLoadbalancerAllConflict{}
}

// WithPayload adds the payload to the delete config loadbalancer all conflict response
func (o *DeleteConfigLoadbalancerAllConflict) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerAllConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer all conflict response
func (o *DeleteConfigLoadbalancerAllConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerAllConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerAllInternalServerErrorCode is the HTTP code returned for type DeleteConfigLoadbalancerAllInternalServerError
const DeleteConfigLoadbalancerAllInternalServerErrorCode int = 500

/*
DeleteConfigLoadbalancerAllInternalServerError Internal service error

swagger:response deleteConfigLoadbalancerAllInternalServerError
*/
type DeleteConfigLoadbalancerAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerAllInternalServerError creates DeleteConfigLoadbalancerAllInternalServerError with default headers values
func NewDeleteConfigLoadbalancerAllInternalServerError() *DeleteConfigLoadbalancerAllInternalServerError {

	return &DeleteConfigLoadbalancerAllInternalServerError{}
}

// WithPayload adds the payload to the delete config loadbalancer all internal server error response
func (o *DeleteConfigLoadbalancerAllInternalServerError) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer all internal server error response
func (o *DeleteConfigLoadbalancerAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerAllServiceUnavailableCode is the HTTP code returned for type DeleteConfigLoadbalancerAllServiceUnavailable
const DeleteConfigLoadbalancerAllServiceUnavailableCode int = 503

/*
DeleteConfigLoadbalancerAllServiceUnavailable Maintanence mode

swagger:response deleteConfigLoadbalancerAllServiceUnavailable
*/
type DeleteConfigLoadbalancerAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerAllServiceUnavailable creates DeleteConfigLoadbalancerAllServiceUnavailable with default headers values
func NewDeleteConfigLoadbalancerAllServiceUnavailable() *DeleteConfigLoadbalancerAllServiceUnavailable {

	return &DeleteConfigLoadbalancerAllServiceUnavailable{}
}

// WithPayload adds the payload to the delete config loadbalancer all service unavailable response
func (o *DeleteConfigLoadbalancerAllServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer all service unavailable response
func (o *DeleteConfigLoadbalancerAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
