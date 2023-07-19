// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// DeleteConfigFirewallURL generates an URL for the delete config firewall operation
type DeleteConfigFirewallURL struct {
	DestinationIP      *string
	MaxDestinationPort *int64
	MaxSourcePort      *int64
	MinDestinationPort *int64
	MinSourcePort      *int64
	PortName           *string
	Preference         *int64
	Protocol           *int64
	SourceIP           *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteConfigFirewallURL) WithBasePath(bp string) *DeleteConfigFirewallURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteConfigFirewallURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteConfigFirewallURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/config/firewall"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/flomesh/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var destinationIPQ string
	if o.DestinationIP != nil {
		destinationIPQ = *o.DestinationIP
	}
	if destinationIPQ != "" {
		qs.Set("destinationIP", destinationIPQ)
	}

	var maxDestinationPortQ string
	if o.MaxDestinationPort != nil {
		maxDestinationPortQ = swag.FormatInt64(*o.MaxDestinationPort)
	}
	if maxDestinationPortQ != "" {
		qs.Set("maxDestinationPort", maxDestinationPortQ)
	}

	var maxSourcePortQ string
	if o.MaxSourcePort != nil {
		maxSourcePortQ = swag.FormatInt64(*o.MaxSourcePort)
	}
	if maxSourcePortQ != "" {
		qs.Set("maxSourcePort", maxSourcePortQ)
	}

	var minDestinationPortQ string
	if o.MinDestinationPort != nil {
		minDestinationPortQ = swag.FormatInt64(*o.MinDestinationPort)
	}
	if minDestinationPortQ != "" {
		qs.Set("minDestinationPort", minDestinationPortQ)
	}

	var minSourcePortQ string
	if o.MinSourcePort != nil {
		minSourcePortQ = swag.FormatInt64(*o.MinSourcePort)
	}
	if minSourcePortQ != "" {
		qs.Set("minSourcePort", minSourcePortQ)
	}

	var portNameQ string
	if o.PortName != nil {
		portNameQ = *o.PortName
	}
	if portNameQ != "" {
		qs.Set("portName", portNameQ)
	}

	var preferenceQ string
	if o.Preference != nil {
		preferenceQ = swag.FormatInt64(*o.Preference)
	}
	if preferenceQ != "" {
		qs.Set("preference", preferenceQ)
	}

	var protocolQ string
	if o.Protocol != nil {
		protocolQ = swag.FormatInt64(*o.Protocol)
	}
	if protocolQ != "" {
		qs.Set("protocol", protocolQ)
	}

	var sourceIPQ string
	if o.SourceIP != nil {
		sourceIPQ = *o.SourceIP
	}
	if sourceIPQ != "" {
		qs.Set("sourceIP", sourceIPQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteConfigFirewallURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteConfigFirewallURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteConfigFirewallURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteConfigFirewallURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteConfigFirewallURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DeleteConfigFirewallURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}