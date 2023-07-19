// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// DeleteConfigRouteDestinationIPNetIPAddressMaskURL generates an URL for the delete config route destination IP net IP address mask operation
type DeleteConfigRouteDestinationIPNetIPAddressMaskURL struct {
	IPAddress string
	Mask      int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskURL) WithBasePath(bp string) *DeleteConfigRouteDestinationIPNetIPAddressMaskURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/config/route/destinationIPNet/{ip_address}/{mask}"

	iPAddress := o.IPAddress
	if iPAddress != "" {
		_path = strings.Replace(_path, "{ip_address}", iPAddress, -1)
	} else {
		return nil, errors.New("ipAddress is required on DeleteConfigRouteDestinationIPNetIPAddressMaskURL")
	}

	mask := swag.FormatInt64(o.Mask)
	if mask != "" {
		_path = strings.Replace(_path, "{mask}", mask, -1)
	} else {
		return nil, errors.New("mask is required on DeleteConfigRouteDestinationIPNetIPAddressMaskURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/flomesh/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteConfigRouteDestinationIPNetIPAddressMaskURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteConfigRouteDestinationIPNetIPAddressMaskURL")
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
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
