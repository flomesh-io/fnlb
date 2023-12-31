// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FileSystemInfoEntry file system info entry
//
// swagger:model FileSystemInfoEntry
type FileSystemInfoEntry struct {

	// size of remain the disk
	Avail string `json:"avail,omitempty"`

	// File system name mounted on this device
	FileSystem string `json:"fileSystem,omitempty"`

	// path of the mounted on
	MountedOn string `json:"mountedOn,omitempty"`

	// Boot ID in the linux
	Size string `json:"size,omitempty"`

	// File type (ex. nfs, ext4..)
	Type string `json:"type,omitempty"`

	// usage per total size
	UsePercent string `json:"usePercent,omitempty"`

	// size of used the disk
	Used string `json:"used,omitempty"`
}

// Validate validates this file system info entry
func (m *FileSystemInfoEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this file system info entry based on context it is used
func (m *FileSystemInfoEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileSystemInfoEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileSystemInfoEntry) UnmarshalBinary(b []byte) error {
	var res FileSystemInfoEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
