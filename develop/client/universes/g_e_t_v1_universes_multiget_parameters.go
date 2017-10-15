// Code generated by go-swagger; DO NOT EDIT.

package universes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGETV1UniversesMultigetParams creates a new GETV1UniversesMultigetParams object
// with the default values initialized.
func NewGETV1UniversesMultigetParams() *GETV1UniversesMultigetParams {
	var ()
	return &GETV1UniversesMultigetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETV1UniversesMultigetParamsWithTimeout creates a new GETV1UniversesMultigetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETV1UniversesMultigetParamsWithTimeout(timeout time.Duration) *GETV1UniversesMultigetParams {
	var ()
	return &GETV1UniversesMultigetParams{

		timeout: timeout,
	}
}

// NewGETV1UniversesMultigetParamsWithContext creates a new GETV1UniversesMultigetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETV1UniversesMultigetParamsWithContext(ctx context.Context) *GETV1UniversesMultigetParams {
	var ()
	return &GETV1UniversesMultigetParams{

		Context: ctx,
	}
}

// NewGETV1UniversesMultigetParamsWithHTTPClient creates a new GETV1UniversesMultigetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETV1UniversesMultigetParamsWithHTTPClient(client *http.Client) *GETV1UniversesMultigetParams {
	var ()
	return &GETV1UniversesMultigetParams{
		HTTPClient: client,
	}
}

/*GETV1UniversesMultigetParams contains all the parameters to send to the API endpoint
for the g e t v1 universes multiget operation typically these are written to a http.Request
*/
type GETV1UniversesMultigetParams struct {

	/*Ids
	  The universe IDs to get. Limit 100.

	*/
	Ids []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the g e t v1 universes multiget params
func (o *GETV1UniversesMultigetParams) WithTimeout(timeout time.Duration) *GETV1UniversesMultigetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g e t v1 universes multiget params
func (o *GETV1UniversesMultigetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g e t v1 universes multiget params
func (o *GETV1UniversesMultigetParams) WithContext(ctx context.Context) *GETV1UniversesMultigetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g e t v1 universes multiget params
func (o *GETV1UniversesMultigetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the g e t v1 universes multiget params
func (o *GETV1UniversesMultigetParams) WithHTTPClient(client *http.Client) *GETV1UniversesMultigetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the g e t v1 universes multiget params
func (o *GETV1UniversesMultigetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the g e t v1 universes multiget params
func (o *GETV1UniversesMultigetParams) WithIds(ids []int64) *GETV1UniversesMultigetParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the g e t v1 universes multiget params
func (o *GETV1UniversesMultigetParams) SetIds(ids []int64) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GETV1UniversesMultigetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	var valuesIds []string
	for _, v := range o.Ids {
		valuesIds = append(valuesIds, swag.FormatInt64(v))
	}

	joinedIds := swag.JoinByFormat(valuesIds, "multi")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
