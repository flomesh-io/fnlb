// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cybwan/fsmxlb/api/models"
)

// GetConfigVlanAllHandlerFunc turns a function with the right signature into a get config vlan all handler
type GetConfigVlanAllHandlerFunc func(GetConfigVlanAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigVlanAllHandlerFunc) Handle(params GetConfigVlanAllParams) middleware.Responder {
	return fn(params)
}

// GetConfigVlanAllHandler interface for that can handle valid get config vlan all params
type GetConfigVlanAllHandler interface {
	Handle(GetConfigVlanAllParams) middleware.Responder
}

// NewGetConfigVlanAll creates a new http.Handler for the get config vlan all operation
func NewGetConfigVlanAll(ctx *middleware.Context, handler GetConfigVlanAllHandler) *GetConfigVlanAll {
	return &GetConfigVlanAll{Context: ctx, Handler: handler}
}

/*
	GetConfigVlanAll swagger:route GET /config/vlan/all getConfigVlanAll

# Get vlan in the device

Get vlan in the device
*/
type GetConfigVlanAll struct {
	Context *middleware.Context
	Handler GetConfigVlanAllHandler
}

func (o *GetConfigVlanAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigVlanAllParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetConfigVlanAllOKBody get config vlan all o k body
//
// swagger:model GetConfigVlanAllOKBody
type GetConfigVlanAllOKBody struct {

	// vlan attr
	VlanAttr []*models.VlanGetEntry `json:"vlanAttr"`
}

// Validate validates this get config vlan all o k body
func (o *GetConfigVlanAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVlanAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigVlanAllOKBody) validateVlanAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.VlanAttr) { // not required
		return nil
	}

	for i := 0; i < len(o.VlanAttr); i++ {
		if swag.IsZero(o.VlanAttr[i]) { // not required
			continue
		}

		if o.VlanAttr[i] != nil {
			if err := o.VlanAttr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigVlanAllOK" + "." + "vlanAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigVlanAllOK" + "." + "vlanAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get config vlan all o k body based on the context it is used
func (o *GetConfigVlanAllOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVlanAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigVlanAllOKBody) contextValidateVlanAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.VlanAttr); i++ {

		if o.VlanAttr[i] != nil {

			if swag.IsZero(o.VlanAttr[i]) { // not required
				return nil
			}

			if err := o.VlanAttr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigVlanAllOK" + "." + "vlanAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigVlanAllOK" + "." + "vlanAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigVlanAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigVlanAllOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigVlanAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
