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

// NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams creates a new DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams object
//
// There are no default values defined in the spec.
func NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams() DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams {

	return DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams{}
}

// DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams contains all the bound params for the delete config vlan vlan ID member if name tagged tagged operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteConfigVlanVlanIDMemberIfNameTaggedTagged
type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Physical port name
	  Required: true
	  In: path
	*/
	IfName string
	/*Tagged status
	  Required: true
	  In: path
	*/
	Tagged bool
	/*12 bit vlan_id
	  Required: true
	  In: path
	*/
	VlanID int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams() beforehand.
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rIfName, rhkIfName, _ := route.Params.GetOK("if_name")
	if err := o.bindIfName(rIfName, rhkIfName, route.Formats); err != nil {
		res = append(res, err)
	}

	rTagged, rhkTagged, _ := route.Params.GetOK("tagged")
	if err := o.bindTagged(rTagged, rhkTagged, route.Formats); err != nil {
		res = append(res, err)
	}

	rVlanID, rhkVlanID, _ := route.Params.GetOK("vlan_id")
	if err := o.bindVlanID(rVlanID, rhkVlanID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIfName binds and validates parameter IfName from path.
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams) bindIfName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.IfName = raw

	return nil
}

// bindTagged binds and validates parameter Tagged from path.
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams) bindTagged(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("tagged", "path", "bool", raw)
	}
	o.Tagged = value

	return nil
}

// bindVlanID binds and validates parameter VlanID from path.
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams) bindVlanID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("vlan_id", "path", "int32", raw)
	}
	o.VlanID = value

	return nil
}
