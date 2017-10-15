// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/auth/client/account_pin"
	"github.com/janberktold/go-roblox/auth/client/passwords"
	"github.com/janberktold/go-roblox/auth/client/saml"
	"github.com/janberktold/go-roblox/auth/client/social_authentication"
	"github.com/janberktold/go-roblox/auth/client/two_step_verification"
	"github.com/janberktold/go-roblox/auth/client/validators"
	"github.com/janberktold/go-roblox/auth/client/xbox"
)

// Default auth HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "auth.roblox.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new auth HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Auth {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new auth HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Auth {
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

// New creates a new auth client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Auth {
	cli := new(Auth)
	cli.Transport = transport

	cli.AccountPin = account_pin.New(transport, formats)

	cli.Passwords = passwords.New(transport, formats)

	cli.Saml = saml.New(transport, formats)

	cli.SocialAuthentication = social_authentication.New(transport, formats)

	cli.TwoStepVerification = two_step_verification.New(transport, formats)

	cli.Validators = validators.New(transport, formats)

	cli.Xbox = xbox.New(transport, formats)

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

// Auth is a client for auth
type Auth struct {
	AccountPin *account_pin.Client

	Passwords *passwords.Client

	Saml *saml.Client

	SocialAuthentication *social_authentication.Client

	TwoStepVerification *two_step_verification.Client

	Validators *validators.Client

	Xbox *xbox.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Auth) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.AccountPin.SetTransport(transport)

	c.Passwords.SetTransport(transport)

	c.Saml.SetTransport(transport)

	c.SocialAuthentication.SetTransport(transport)

	c.TwoStepVerification.SetTransport(transport)

	c.Validators.SetTransport(transport)

	c.Xbox.SetTransport(transport)

}