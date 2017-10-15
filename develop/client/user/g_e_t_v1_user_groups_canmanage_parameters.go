// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGETV1UserGroupsCanmanageParams creates a new GETV1UserGroupsCanmanageParams object
// with the default values initialized.
func NewGETV1UserGroupsCanmanageParams() *GETV1UserGroupsCanmanageParams {

	return &GETV1UserGroupsCanmanageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETV1UserGroupsCanmanageParamsWithTimeout creates a new GETV1UserGroupsCanmanageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETV1UserGroupsCanmanageParamsWithTimeout(timeout time.Duration) *GETV1UserGroupsCanmanageParams {

	return &GETV1UserGroupsCanmanageParams{

		timeout: timeout,
	}
}

// NewGETV1UserGroupsCanmanageParamsWithContext creates a new GETV1UserGroupsCanmanageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETV1UserGroupsCanmanageParamsWithContext(ctx context.Context) *GETV1UserGroupsCanmanageParams {

	return &GETV1UserGroupsCanmanageParams{

		Context: ctx,
	}
}

// NewGETV1UserGroupsCanmanageParamsWithHTTPClient creates a new GETV1UserGroupsCanmanageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETV1UserGroupsCanmanageParamsWithHTTPClient(client *http.Client) *GETV1UserGroupsCanmanageParams {

	return &GETV1UserGroupsCanmanageParams{
		HTTPClient: client,
	}
}

/*GETV1UserGroupsCanmanageParams contains all the parameters to send to the API endpoint
for the g e t v1 user groups canmanage operation typically these are written to a http.Request
*/
type GETV1UserGroupsCanmanageParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the g e t v1 user groups canmanage params
func (o *GETV1UserGroupsCanmanageParams) WithTimeout(timeout time.Duration) *GETV1UserGroupsCanmanageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g e t v1 user groups canmanage params
func (o *GETV1UserGroupsCanmanageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g e t v1 user groups canmanage params
func (o *GETV1UserGroupsCanmanageParams) WithContext(ctx context.Context) *GETV1UserGroupsCanmanageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g e t v1 user groups canmanage params
func (o *GETV1UserGroupsCanmanageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the g e t v1 user groups canmanage params
func (o *GETV1UserGroupsCanmanageParams) WithHTTPClient(client *http.Client) *GETV1UserGroupsCanmanageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the g e t v1 user groups canmanage params
func (o *GETV1UserGroupsCanmanageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GETV1UserGroupsCanmanageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}