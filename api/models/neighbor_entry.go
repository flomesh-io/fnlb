// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NeighborEntry neighbor entry
//
// swagger:model NeighborEntry
type NeighborEntry struct {

	// Name of the interface device to which you want to add neighbor
	Dev string `json:"dev,omitempty"`

	// IP address to neighbor
	IPAddress string `json:"ipAddress,omitempty"`

	// MAC address to neighbor
	MacAddress string `json:"macAddress,omitempty"`
}

// Validate validates this neighbor entry
func (m *NeighborEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this neighbor entry based on context it is used
func (m *NeighborEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NeighborEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeighborEntry) UnmarshalBinary(b []byte) error {
	var res NeighborEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
