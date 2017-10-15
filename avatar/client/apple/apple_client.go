// Code generated by go-swagger; DO NOT EDIT.

package apple

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new apple API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for apple API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
POSTV1ApplePurchase performs a purchase and grant desired product
*/
func (a *Client) POSTV1ApplePurchase(params *POSTV1ApplePurchaseParams) (*POSTV1ApplePurchaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTV1ApplePurchaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST_v1_apple_purchase",
		Method:             "POST",
		PathPattern:        "/v1/apple/purchase",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTV1ApplePurchaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTV1ApplePurchaseOK), nil

}

/*
POSTV1AppleValidate validates a product Id before making a purchase
*/
func (a *Client) POSTV1AppleValidate(params *POSTV1AppleValidateParams) (*POSTV1AppleValidateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTV1AppleValidateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST_v1_apple_validate",
		Method:             "POST",
		PathPattern:        "/v1/apple/validate",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTV1AppleValidateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTV1AppleValidateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}