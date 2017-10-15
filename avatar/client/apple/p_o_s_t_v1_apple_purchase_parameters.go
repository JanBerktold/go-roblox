// Code generated by go-swagger; DO NOT EDIT.

package apple

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

	"github.com/janberktold/go-roblox/avatar/models"
)

// NewPOSTV1ApplePurchaseParams creates a new POSTV1ApplePurchaseParams object
// with the default values initialized.
func NewPOSTV1ApplePurchaseParams() *POSTV1ApplePurchaseParams {
	var ()
	return &POSTV1ApplePurchaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTV1ApplePurchaseParamsWithTimeout creates a new POSTV1ApplePurchaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPOSTV1ApplePurchaseParamsWithTimeout(timeout time.Duration) *POSTV1ApplePurchaseParams {
	var ()
	return &POSTV1ApplePurchaseParams{

		timeout: timeout,
	}
}

// NewPOSTV1ApplePurchaseParamsWithContext creates a new POSTV1ApplePurchaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewPOSTV1ApplePurchaseParamsWithContext(ctx context.Context) *POSTV1ApplePurchaseParams {
	var ()
	return &POSTV1ApplePurchaseParams{

		Context: ctx,
	}
}

// NewPOSTV1ApplePurchaseParamsWithHTTPClient creates a new POSTV1ApplePurchaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPOSTV1ApplePurchaseParamsWithHTTPClient(client *http.Client) *POSTV1ApplePurchaseParams {
	var ()
	return &POSTV1ApplePurchaseParams{
		HTTPClient: client,
	}
}

/*POSTV1ApplePurchaseParams contains all the parameters to send to the API endpoint
for the p o s t v1 apple purchase operation typically these are written to a http.Request
*/
type POSTV1ApplePurchaseParams struct {

	/*AppleStorePurchaseModel
	  Apple Purchase Model

	*/
	AppleStorePurchaseModel *models.AppleStorePurchaseModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the p o s t v1 apple purchase params
func (o *POSTV1ApplePurchaseParams) WithTimeout(timeout time.Duration) *POSTV1ApplePurchaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p o s t v1 apple purchase params
func (o *POSTV1ApplePurchaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p o s t v1 apple purchase params
func (o *POSTV1ApplePurchaseParams) WithContext(ctx context.Context) *POSTV1ApplePurchaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p o s t v1 apple purchase params
func (o *POSTV1ApplePurchaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the p o s t v1 apple purchase params
func (o *POSTV1ApplePurchaseParams) WithHTTPClient(client *http.Client) *POSTV1ApplePurchaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the p o s t v1 apple purchase params
func (o *POSTV1ApplePurchaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppleStorePurchaseModel adds the appleStorePurchaseModel to the p o s t v1 apple purchase params
func (o *POSTV1ApplePurchaseParams) WithAppleStorePurchaseModel(appleStorePurchaseModel *models.AppleStorePurchaseModel) *POSTV1ApplePurchaseParams {
	o.SetAppleStorePurchaseModel(appleStorePurchaseModel)
	return o
}

// SetAppleStorePurchaseModel adds the appleStorePurchaseModel to the p o s t v1 apple purchase params
func (o *POSTV1ApplePurchaseParams) SetAppleStorePurchaseModel(appleStorePurchaseModel *models.AppleStorePurchaseModel) {
	o.AppleStorePurchaseModel = appleStorePurchaseModel
}

// WriteToRequest writes these params to a swagger request
func (o *POSTV1ApplePurchaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppleStorePurchaseModel != nil {
		if err := r.SetBodyParam(o.AppleStorePurchaseModel); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
