// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/client/status"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
)

// Default net box API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:8001"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new net box API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *NetBoxAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new net box API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *NetBoxAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new net box API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *NetBoxAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(NetBoxAPI)
	cli.Transport = transport
	cli.Circuits = circuits.New(transport, formats)
	cli.Dcim = dcim.New(transport, formats)
	cli.Extras = extras.New(transport, formats)
	cli.Ipam = ipam.New(transport, formats)
	cli.Status = status.New(transport, formats)
	cli.Tenancy = tenancy.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.Virtualization = virtualization.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// NetBoxAPI is a client for net box API
type NetBoxAPI struct {
	Circuits circuits.ClientService

	Dcim dcim.ClientService

	Extras extras.ClientService

	Ipam ipam.ClientService

	Status status.ClientService

	Tenancy tenancy.ClientService

	Users users.ClientService

	Virtualization virtualization.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *NetBoxAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Circuits.SetTransport(transport)
	c.Dcim.SetTransport(transport)
	c.Extras.SetTransport(transport)
	c.Ipam.SetTransport(transport)
	c.Status.SetTransport(transport)
	c.Tenancy.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.Virtualization.SetTransport(transport)
}
