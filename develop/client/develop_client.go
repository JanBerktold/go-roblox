// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/develop/client/game_templates"
	"github.com/janberktold/go-roblox/develop/client/groups"
	"github.com/janberktold/go-roblox/develop/client/places"
	"github.com/janberktold/go-roblox/develop/client/team_create"
	"github.com/janberktold/go-roblox/develop/client/universe_settings"
	"github.com/janberktold/go-roblox/develop/client/universes"
	"github.com/janberktold/go-roblox/develop/client/user"
)

// Default develop HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "develop.roblox.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new develop HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Develop {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new develop HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Develop {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new develop client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Develop {
	cli := new(Develop)
	cli.Transport = transport

	cli.GameTemplates = game_templates.New(transport, formats)

	cli.Groups = groups.New(transport, formats)

	cli.Places = places.New(transport, formats)

	cli.TeamCreate = team_create.New(transport, formats)

	cli.UniverseSettings = universe_settings.New(transport, formats)

	cli.Universes = universes.New(transport, formats)

	cli.User = user.New(transport, formats)

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

// Develop is a client for develop
type Develop struct {
	GameTemplates *game_templates.Client

	Groups *groups.Client

	Places *places.Client

	TeamCreate *team_create.Client

	UniverseSettings *universe_settings.Client

	Universes *universes.Client

	User *user.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Develop) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.GameTemplates.SetTransport(transport)

	c.Groups.SetTransport(transport)

	c.Places.SetTransport(transport)

	c.TeamCreate.SetTransport(transport)

	c.UniverseSettings.SetTransport(transport)

	c.Universes.SetTransport(transport)

	c.User.SetTransport(transport)

}
