// Code generated by go-swagger; DO NOT EDIT.

package places

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

	"github.com/janberktold/go-roblox/develop/models"
)

// NewPATCHV1PlacesPlaceIDParams creates a new PATCHV1PlacesPlaceIDParams object
// with the default values initialized.
func NewPATCHV1PlacesPlaceIDParams() *PATCHV1PlacesPlaceIDParams {
	var ()
	return &PATCHV1PlacesPlaceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPATCHV1PlacesPlaceIDParamsWithTimeout creates a new PATCHV1PlacesPlaceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPATCHV1PlacesPlaceIDParamsWithTimeout(timeout time.Duration) *PATCHV1PlacesPlaceIDParams {
	var ()
	return &PATCHV1PlacesPlaceIDParams{

		timeout: timeout,
	}
}

// NewPATCHV1PlacesPlaceIDParamsWithContext creates a new PATCHV1PlacesPlaceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPATCHV1PlacesPlaceIDParamsWithContext(ctx context.Context) *PATCHV1PlacesPlaceIDParams {
	var ()
	return &PATCHV1PlacesPlaceIDParams{

		Context: ctx,
	}
}

// NewPATCHV1PlacesPlaceIDParamsWithHTTPClient creates a new PATCHV1PlacesPlaceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPATCHV1PlacesPlaceIDParamsWithHTTPClient(client *http.Client) *PATCHV1PlacesPlaceIDParams {
	var ()
	return &PATCHV1PlacesPlaceIDParams{
		HTTPClient: client,
	}
}

/*PATCHV1PlacesPlaceIDParams contains all the parameters to send to the API endpoint
for the p a t c h v1 places place Id operation typically these are written to a http.Request
*/
type PATCHV1PlacesPlaceIDParams struct {

	/*Configuration*/
	Configuration *models.PlaceConfigurationModel
	/*PlaceID
	  The place id for the place to be updated.

	*/
	PlaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the p a t c h v1 places place Id params
func (o *PATCHV1PlacesPlaceIDParams) WithTimeout(timeout time.Duration) *PATCHV1PlacesPlaceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p a t c h v1 places place Id params
func (o *PATCHV1PlacesPlaceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p a t c h v1 places place Id params
func (o *PATCHV1PlacesPlaceIDParams) WithContext(ctx context.Context) *PATCHV1PlacesPlaceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p a t c h v1 places place Id params
func (o *PATCHV1PlacesPlaceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the p a t c h v1 places place Id params
func (o *PATCHV1PlacesPlaceIDParams) WithHTTPClient(client *http.Client) *PATCHV1PlacesPlaceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the p a t c h v1 places place Id params
func (o *PATCHV1PlacesPlaceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfiguration adds the configuration to the p a t c h v1 places place Id params
func (o *PATCHV1PlacesPlaceIDParams) WithConfiguration(configuration *models.PlaceConfigurationModel) *PATCHV1PlacesPlaceIDParams {
	o.SetConfiguration(configuration)
	return o
}

// SetConfiguration adds the configuration to the p a t c h v1 places place Id params
func (o *PATCHV1PlacesPlaceIDParams) SetConfiguration(configuration *models.PlaceConfigurationModel) {
	o.Configuration = configuration
}

// WithPlaceID adds the placeID to the p a t c h v1 places place Id params
func (o *PATCHV1PlacesPlaceIDParams) WithPlaceID(placeID int64) *PATCHV1PlacesPlaceIDParams {
	o.SetPlaceID(placeID)
	return o
}

// SetPlaceID adds the placeId to the p a t c h v1 places place Id params
func (o *PATCHV1PlacesPlaceIDParams) SetPlaceID(placeID int64) {
	o.PlaceID = placeID
}

// WriteToRequest writes these params to a swagger request
func (o *PATCHV1PlacesPlaceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Configuration != nil {
		if err := r.SetBodyParam(o.Configuration); err != nil {
			return err
		}
	}

	// path param placeId
	if err := r.SetPathParam("placeId", swag.FormatInt64(o.PlaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}