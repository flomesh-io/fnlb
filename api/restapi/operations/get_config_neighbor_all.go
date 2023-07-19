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

// GetConfigNeighborAllHandlerFunc turns a function with the right signature into a get config neighbor all handler
type GetConfigNeighborAllHandlerFunc func(GetConfigNeighborAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigNeighborAllHandlerFunc) Handle(params GetConfigNeighborAllParams) middleware.Responder {
	return fn(params)
}

// GetConfigNeighborAllHandler interface for that can handle valid get config neighbor all params
type GetConfigNeighborAllHandler interface {
	Handle(GetConfigNeighborAllParams) middleware.Responder
}

// NewGetConfigNeighborAll creates a new http.Handler for the get config neighbor all operation
func NewGetConfigNeighborAll(ctx *middleware.Context, handler GetConfigNeighborAllHandler) *GetConfigNeighborAll {
	return &GetConfigNeighborAll{Context: ctx, Handler: handler}
}

/*
	GetConfigNeighborAll swagger:route GET /config/neighbor/all getConfigNeighborAll

Get IPv4 neighbor in the device(interface)

Get IPv4 neighbor in the device(interface)
*/
type GetConfigNeighborAll struct {
	Context *middleware.Context
	Handler GetConfigNeighborAllHandler
}

func (o *GetConfigNeighborAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigNeighborAllParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetConfigNeighborAllOKBody get config neighbor all o k body
//
// swagger:model GetConfigNeighborAllOKBody
type GetConfigNeighborAllOKBody struct {

	// neighbor attr
	NeighborAttr []*models.NeighborEntry `json:"neighborAttr"`
}

// Validate validates this get config neighbor all o k body
func (o *GetConfigNeighborAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNeighborAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigNeighborAllOKBody) validateNeighborAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.NeighborAttr) { // not required
		return nil
	}

	for i := 0; i < len(o.NeighborAttr); i++ {
		if swag.IsZero(o.NeighborAttr[i]) { // not required
			continue
		}

		if o.NeighborAttr[i] != nil {
			if err := o.NeighborAttr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigNeighborAllOK" + "." + "neighborAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigNeighborAllOK" + "." + "neighborAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get config neighbor all o k body based on the context it is used
func (o *GetConfigNeighborAllOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNeighborAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigNeighborAllOKBody) contextValidateNeighborAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.NeighborAttr); i++ {

		if o.NeighborAttr[i] != nil {

			if swag.IsZero(o.NeighborAttr[i]) { // not required
				return nil
			}

			if err := o.NeighborAttr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigNeighborAllOK" + "." + "neighborAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigNeighborAllOK" + "." + "neighborAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigNeighborAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigNeighborAllOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigNeighborAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
