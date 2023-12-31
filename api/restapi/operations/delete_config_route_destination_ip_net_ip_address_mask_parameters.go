// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteConfigRouteDestinationIPNetIPAddressMaskParams creates a new DeleteConfigRouteDestinationIPNetIPAddressMaskParams object
//
// There are no default values defined in the spec.
func NewDeleteConfigRouteDestinationIPNetIPAddressMaskParams() DeleteConfigRouteDestinationIPNetIPAddressMaskParams {

	return DeleteConfigRouteDestinationIPNetIPAddressMaskParams{}
}

// DeleteConfigRouteDestinationIPNetIPAddressMaskParams contains all the bound params for the delete config route destination IP net IP address mask operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteConfigRouteDestinationIPNetIPAddressMask
type DeleteConfigRouteDestinationIPNetIPAddressMaskParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Attributes for destinaion route address
	  Required: true
	  In: path
	*/
	IPAddress string
	/*Attributes for destination route
	  Required: true
	  In: path
	*/
	Mask int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteConfigRouteDestinationIPNetIPAddressMaskParams() beforehand.
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rIPAddress, rhkIPAddress, _ := route.Params.GetOK("ip_address")
	if err := o.bindIPAddress(rIPAddress, rhkIPAddress, route.Formats); err != nil {
		res = append(res, err)
	}

	rMask, rhkMask, _ := route.Params.GetOK("mask")
	if err := o.bindMask(rMask, rhkMask, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIPAddress binds and validates parameter IPAddress from path.
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskParams) bindIPAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.IPAddress = raw

	return nil
}

// bindMask binds and validates parameter Mask from path.
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskParams) bindMask(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("mask", "path", "int64", raw)
	}
	o.Mask = value

	return nil
}
