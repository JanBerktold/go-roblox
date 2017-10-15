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

// NewPOSTV1GooglePurchaseParams creates a new POSTV1GooglePurchaseParams object
// with the default values initialized.
func NewPOSTV1GooglePurchaseParams() *POSTV1GooglePurchaseParams {
	var ()
	return &POSTV1GooglePurchaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTV1GooglePurchaseParamsWithTimeout creates a new POSTV1GooglePurchaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPOSTV1GooglePurchaseParamsWithTimeout(timeout time.Duration) *POSTV1GooglePurchaseParams {
	var ()
	return &POSTV1GooglePurchaseParams{

		timeout: timeout,
	}
}

// NewPOSTV1GooglePurchaseParamsWithContext creates a new POSTV1GooglePurchaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewPOSTV1GooglePurchaseParamsWithContext(ctx context.Context) *POSTV1GooglePurchaseParams {
	var ()
	return &POSTV1GooglePurchaseParams{

		Context: ctx,
	}
}

// NewPOSTV1GooglePurchaseParamsWithHTTPClient creates a new POSTV1GooglePurchaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPOSTV1GooglePurchaseParamsWithHTTPClient(client *http.Client) *POSTV1GooglePurchaseParams {
	var ()
	return &POSTV1GooglePurchaseParams{
		HTTPClient: client,
	}
}

/*POSTV1GooglePurchaseParams contains all the parameters to send to the API endpoint
for the p o s t v1 google purchase operation typically these are written to a http.Request
*/
type POSTV1GooglePurchaseParams struct {

	/*PurchaseModel*/
	PurchaseModel *models.GooglePlayPurchaseModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the p o s t v1 google purchase params
func (o *POSTV1GooglePurchaseParams) WithTimeout(timeout time.Duration) *POSTV1GooglePurchaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p o s t v1 google purchase params
func (o *POSTV1GooglePurchaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p o s t v1 google purchase params
func (o *POSTV1GooglePurchaseParams) WithContext(ctx context.Context) *POSTV1GooglePurchaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p o s t v1 google purchase params
func (o *POSTV1GooglePurchaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the p o s t v1 google purchase params
func (o *POSTV1GooglePurchaseParams) WithHTTPClient(client *http.Client) *POSTV1GooglePurchaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the p o s t v1 google purchase params
func (o *POSTV1GooglePurchaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPurchaseModel adds the purchaseModel to the p o s t v1 google purchase params
func (o *POSTV1GooglePurchaseParams) WithPurchaseModel(purchaseModel *models.GooglePlayPurchaseModel) *POSTV1GooglePurchaseParams {
	o.SetPurchaseModel(purchaseModel)
	return o
}

// SetPurchaseModel adds the purchaseModel to the p o s t v1 google purchase params
func (o *POSTV1GooglePurchaseParams) SetPurchaseModel(purchaseModel *models.GooglePlayPurchaseModel) {
	o.PurchaseModel = purchaseModel
}

// WriteToRequest writes these params to a swagger request
func (o *POSTV1GooglePurchaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PurchaseModel != nil {
		if err := r.SetBodyParam(o.PurchaseModel); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
