// Code generated by go-swagger; DO NOT EDIT.

package wall

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

// NewGETV1GroupsGroupIDWallPostsParams creates a new GETV1GroupsGroupIDWallPostsParams object
// with the default values initialized.
func NewGETV1GroupsGroupIDWallPostsParams() *GETV1GroupsGroupIDWallPostsParams {
	var (
		limitDefault     = int32(10)
		sortOrderDefault = string("Asc")
	)
	return &GETV1GroupsGroupIDWallPostsParams{
		Limit:     &limitDefault,
		SortOrder: &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGETV1GroupsGroupIDWallPostsParamsWithTimeout creates a new GETV1GroupsGroupIDWallPostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETV1GroupsGroupIDWallPostsParamsWithTimeout(timeout time.Duration) *GETV1GroupsGroupIDWallPostsParams {
	var (
		limitDefault     = int32(10)
		sortOrderDefault = string("Asc")
	)
	return &GETV1GroupsGroupIDWallPostsParams{
		Limit:     &limitDefault,
		SortOrder: &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGETV1GroupsGroupIDWallPostsParamsWithContext creates a new GETV1GroupsGroupIDWallPostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETV1GroupsGroupIDWallPostsParamsWithContext(ctx context.Context) *GETV1GroupsGroupIDWallPostsParams {
	var (
		limitDefault     = int32(10)
		sortOrderDefault = string("Asc")
	)
	return &GETV1GroupsGroupIDWallPostsParams{
		Limit:     &limitDefault,
		SortOrder: &sortOrderDefault,

		Context: ctx,
	}
}

// NewGETV1GroupsGroupIDWallPostsParamsWithHTTPClient creates a new GETV1GroupsGroupIDWallPostsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETV1GroupsGroupIDWallPostsParamsWithHTTPClient(client *http.Client) *GETV1GroupsGroupIDWallPostsParams {
	var (
		limitDefault     = int32(10)
		sortOrderDefault = string("Asc")
	)
	return &GETV1GroupsGroupIDWallPostsParams{
		Limit:      &limitDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GETV1GroupsGroupIDWallPostsParams contains all the parameters to send to the API endpoint
for the g e t v1 groups group Id wall posts operation typically these are written to a http.Request
*/
type GETV1GroupsGroupIDWallPostsParams struct {

	/*Cursor
	  The paging cursor for the previous or next page.

	*/
	Cursor *string
	/*GroupID
	  The group id.

	*/
	GroupID int32
	/*Limit
	  The amount of results per request.

	*/
	Limit *int32
	/*SortOrder
	  Sorted by group wall post Id

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) WithTimeout(timeout time.Duration) *GETV1GroupsGroupIDWallPostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) WithContext(ctx context.Context) *GETV1GroupsGroupIDWallPostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) WithHTTPClient(client *http.Client) *GETV1GroupsGroupIDWallPostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) WithCursor(cursor *string) *GETV1GroupsGroupIDWallPostsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithGroupID adds the groupID to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) WithGroupID(groupID int32) *GETV1GroupsGroupIDWallPostsParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) SetGroupID(groupID int32) {
	o.GroupID = groupID
}

// WithLimit adds the limit to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) WithLimit(limit *int32) *GETV1GroupsGroupIDWallPostsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithSortOrder adds the sortOrder to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) WithSortOrder(sortOrder *string) *GETV1GroupsGroupIDWallPostsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the g e t v1 groups group Id wall posts params
func (o *GETV1GroupsGroupIDWallPostsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GETV1GroupsGroupIDWallPostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	// path param groupId
	if err := r.SetPathParam("groupId", swag.FormatInt32(o.GroupID)); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string
		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {
			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
