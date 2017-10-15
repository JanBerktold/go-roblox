// Code generated by go-swagger; DO NOT EDIT.

package account_pin

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

	"github.com/janberktold/go-roblox/auth/models"
)

// NewPATCHV1AccountPinParams creates a new PATCHV1AccountPinParams object
// with the default values initialized.
func NewPATCHV1AccountPinParams() *PATCHV1AccountPinParams {
	var ()
	return &PATCHV1AccountPinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPATCHV1AccountPinParamsWithTimeout creates a new PATCHV1AccountPinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPATCHV1AccountPinParamsWithTimeout(timeout time.Duration) *PATCHV1AccountPinParams {
	var ()
	return &PATCHV1AccountPinParams{

		timeout: timeout,
	}
}

// NewPATCHV1AccountPinParamsWithContext creates a new PATCHV1AccountPinParams object
// with the default values initialized, and the ability to set a context for a request
func NewPATCHV1AccountPinParamsWithContext(ctx context.Context) *PATCHV1AccountPinParams {
	var ()
	return &PATCHV1AccountPinParams{

		Context: ctx,
	}
}

// NewPATCHV1AccountPinParamsWithHTTPClient creates a new PATCHV1AccountPinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPATCHV1AccountPinParamsWithHTTPClient(client *http.Client) *PATCHV1AccountPinParams {
	var ()
	return &PATCHV1AccountPinParams{
		HTTPClient: client,
	}
}

/*PATCHV1AccountPinParams contains all the parameters to send to the API endpoint
for the p a t c h v1 account pin operation typically these are written to a http.Request
*/
type PATCHV1AccountPinParams struct {

	/*RequestBody
	  The request body.

	*/
	RequestBody *models.AccountPinRequestModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the p a t c h v1 account pin params
func (o *PATCHV1AccountPinParams) WithTimeout(timeout time.Duration) *PATCHV1AccountPinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p a t c h v1 account pin params
func (o *PATCHV1AccountPinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p a t c h v1 account pin params
func (o *PATCHV1AccountPinParams) WithContext(ctx context.Context) *PATCHV1AccountPinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p a t c h v1 account pin params
func (o *PATCHV1AccountPinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the p a t c h v1 account pin params
func (o *PATCHV1AccountPinParams) WithHTTPClient(client *http.Client) *PATCHV1AccountPinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the p a t c h v1 account pin params
func (o *PATCHV1AccountPinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the p a t c h v1 account pin params
func (o *PATCHV1AccountPinParams) WithRequestBody(requestBody *models.AccountPinRequestModel) *PATCHV1AccountPinParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the p a t c h v1 account pin params
func (o *PATCHV1AccountPinParams) SetRequestBody(requestBody *models.AccountPinRequestModel) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *PATCHV1AccountPinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RequestBody != nil {
		if err := r.SetBodyParam(o.RequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
