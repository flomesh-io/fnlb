// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceInfoEntry device info entry
//
// swagger:model DeviceInfoEntry
type DeviceInfoEntry struct {

	// Operation System of the device
	OS string `json:"OS,omitempty"`

	// CPU architecture of the device
	Architecture string `json:"architecture,omitempty"`

	// Boot ID in the linux
	BootID string `json:"bootID,omitempty"`

	// Device host name
	HostName string `json:"hostName,omitempty"`

	// Kernel version of the device
	Kernel string `json:"kernel,omitempty"`

	// Device machine ID
	MachineID string `json:"machineID,omitempty"`

	// system uptime
	Uptime string `json:"uptime,omitempty"`
}

// Validate validates this device info entry
func (m *DeviceInfoEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this device info entry based on context it is used
func (m *DeviceInfoEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceInfoEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceInfoEntry) UnmarshalBinary(b []byte) error {
	var res DeviceInfoEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
