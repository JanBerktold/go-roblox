// Code generated by go-swagger; DO NOT EDIT.

package saml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new saml API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for saml API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GETV1SamlMetadata gets the s a m l2 metadata XML

This endpoint is exceptional for SAML which is XML-based.
Most Apis should not return XML.
*/
func (a *Client) GETV1SamlMetadata(params *GETV1SamlMetadataParams) (*GETV1SamlMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETV1SamlMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GET_v1_saml_metadata",
		Method:             "GET",
		PathPattern:        "/v1/saml/metadata",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETV1SamlMetadataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GETV1SamlMetadataOK), nil

}

/*
POSTV1SamlLogin authenticates a user for a service through s a m l2

This endpoint is very exceptional.
It can return a redirect response to www, or a SAML2 response.
Most Apis should not return anything except for Json.
*/
func (a *Client) POSTV1SamlLogin(params *POSTV1SamlLoginParams) (*POSTV1SamlLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTV1SamlLoginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST_v1_saml_login",
		Method:             "POST",
		PathPattern:        "/v1/saml/login",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTV1SamlLoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTV1SamlLoginOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
