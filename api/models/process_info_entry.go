// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProcessInfoEntry process info entry
//
// swagger:model ProcessInfoEntry
type ProcessInfoEntry struct {

	// CPU usage of the process
	CPUUsage string `json:"CPUUsage,omitempty"`

	// Memory usage of the process
	MemoryUsage string `json:"MemoryUsage,omitempty"`

	// process command
	Command string `json:"command,omitempty"`

	// process nice value
	Nice string `json:"nice,omitempty"`

	// process ID
	Pid string `json:"pid,omitempty"`

	// process priority
	Priority string `json:"priority,omitempty"`

	// Physical memory usage
	ResidentSize string `json:"residentSize,omitempty"`

	// Shared memory usage
	SharedMemory string `json:"sharedMemory,omitempty"`

	// process status
	Status string `json:"status,omitempty"`

	// Executation time
	Time string `json:"time,omitempty"`

	// User name that start the process
	User string `json:"user,omitempty"`

	// virtual memory usage
	VirtMemory string `json:"virtMemory,omitempty"`
}

// Validate validates this process info entry
func (m *ProcessInfoEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this process info entry based on context it is used
func (m *ProcessInfoEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProcessInfoEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessInfoEntry) UnmarshalBinary(b []byte) error {
	var res ProcessInfoEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
