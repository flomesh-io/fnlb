// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/cybwan/fsmxlb/api/models"
)

// NewPostConfigNeighborParams creates a new PostConfigNeighborParams object
//
// There are no default values defined in the spec.
func NewPostConfigNeighborParams() PostConfigNeighborParams {

	return PostConfigNeighborParams{}
}

// PostConfigNeighborParams contains all the bound params for the post config neighbor operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostConfigNeighbor
type PostConfigNeighborParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Attributes for IPv4 address
	  Required: true
	  In: body
	*/
	Attr *models.NeighborEntry
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostConfigNeighborParams() beforehand.
func (o *PostConfigNeighborParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.NeighborEntry
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("attr", "body", ""))
			} else {
				res = append(res, errors.NewParseError("attr", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Attr = &body
			}
		}
	} else {
		res = append(res, errors.Required("attr", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}