// Code generated by go-swagger; DO NOT EDIT.

package amazon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new amazon API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for amazon API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
POSTV1AmazonPurchase performs a purchase and grant desired product
*/
func (a *Client) POSTV1AmazonPurchase(params *POSTV1AmazonPurchaseParams) (*POSTV1AmazonPurchaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTV1AmazonPurchaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST_v1_amazon_purchase",
		Method:             "POST",
		PathPattern:        "/v1/amazon/purchase",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTV1AmazonPurchaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTV1AmazonPurchaseOK), nil

}

/*
POSTV1AmazonValidate validates a product Id before making a purchase
*/
func (a *Client) POSTV1AmazonValidate(params *POSTV1AmazonValidateParams) (*POSTV1AmazonValidateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTV1AmazonValidateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST_v1_amazon_validate",
		Method:             "POST",
		PathPattern:        "/v1/amazon/validate",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTV1AmazonValidateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTV1AmazonValidateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}