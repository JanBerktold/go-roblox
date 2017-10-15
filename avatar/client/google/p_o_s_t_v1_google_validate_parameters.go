// Code generated by go-swagger; DO NOT EDIT.

package google

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

// NewPOSTV1GoogleValidateParams creates a new POSTV1GoogleValidateParams object
// with the default values initialized.
func NewPOSTV1GoogleValidateParams() *POSTV1GoogleValidateParams {
	var ()
	return &POSTV1GoogleValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTV1GoogleValidateParamsWithTimeout creates a new POSTV1GoogleValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPOSTV1GoogleValidateParamsWithTimeout(timeout time.Duration) *POSTV1GoogleValidateParams {
	var ()
	return &POSTV1GoogleValidateParams{

		timeout: timeout,
	}
}

// NewPOSTV1GoogleValidateParamsWithContext creates a new POSTV1GoogleValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPOSTV1GoogleValidateParamsWithContext(ctx context.Context) *POSTV1GoogleValidateParams {
	var ()
	return &POSTV1GoogleValidateParams{

		Context: ctx,
	}
}

// NewPOSTV1GoogleValidateParamsWithHTTPClient creates a new POSTV1GoogleValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPOSTV1GoogleValidateParamsWithHTTPClient(client *http.Client) *POSTV1GoogleValidateParams {
	var ()
	return &POSTV1GoogleValidateParams{
		HTTPClient: client,
	}
}

/*POSTV1GoogleValidateParams contains all the parameters to send to the API endpoint
for the p o s t v1 google validate operation typically these are written to a http.Request
*/
type POSTV1GoogleValidateParams struct {

	/*ValidateModel
	  Product Id for example: com.roblox.client.bc1month, com.roblox.client.tbc1month, ...

	*/
	ValidateModel *models.ValidateModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the p o s t v1 google validate params
func (o *POSTV1GoogleValidateParams) WithTimeout(timeout time.Duration) *POSTV1GoogleValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p o s t v1 google validate params
func (o *POSTV1GoogleValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p o s t v1 google validate params
func (o *POSTV1GoogleValidateParams) WithContext(ctx context.Context) *POSTV1GoogleValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p o s t v1 google validate params
func (o *POSTV1GoogleValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the p o s t v1 google validate params
func (o *POSTV1GoogleValidateParams) WithHTTPClient(client *http.Client) *POSTV1GoogleValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the p o s t v1 google validate params
func (o *POSTV1GoogleValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidateModel adds the validateModel to the p o s t v1 google validate params
func (o *POSTV1GoogleValidateParams) WithValidateModel(validateModel *models.ValidateModel) *POSTV1GoogleValidateParams {
	o.SetValidateModel(validateModel)
	return o
}

// SetValidateModel adds the validateModel to the p o s t v1 google validate params
func (o *POSTV1GoogleValidateParams) SetValidateModel(validateModel *models.ValidateModel) {
	o.ValidateModel = validateModel
}

// WriteToRequest writes these params to a swagger request
func (o *POSTV1GoogleValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ValidateModel != nil {
		if err := r.SetBodyParam(o.ValidateModel); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
