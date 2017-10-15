// Code generated by go-swagger; DO NOT EDIT.

package membership

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

// NewGETV1GroupsGroupIDRolesRoleSetIDUsersParams creates a new GETV1GroupsGroupIDRolesRoleSetIDUsersParams object
// with the default values initialized.
func NewGETV1GroupsGroupIDRolesRoleSetIDUsersParams() *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	var (
		limitDefault     = int32(10)
		sortOrderDefault = string("Asc")
	)
	return &GETV1GroupsGroupIDRolesRoleSetIDUsersParams{
		Limit:     &limitDefault,
		SortOrder: &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGETV1GroupsGroupIDRolesRoleSetIDUsersParamsWithTimeout creates a new GETV1GroupsGroupIDRolesRoleSetIDUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETV1GroupsGroupIDRolesRoleSetIDUsersParamsWithTimeout(timeout time.Duration) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	var (
		limitDefault     = int32(10)
		sortOrderDefault = string("Asc")
	)
	return &GETV1GroupsGroupIDRolesRoleSetIDUsersParams{
		Limit:     &limitDefault,
		SortOrder: &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGETV1GroupsGroupIDRolesRoleSetIDUsersParamsWithContext creates a new GETV1GroupsGroupIDRolesRoleSetIDUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETV1GroupsGroupIDRolesRoleSetIDUsersParamsWithContext(ctx context.Context) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	var (
		limitDefault     = int32(10)
		sortOrderDefault = string("Asc")
	)
	return &GETV1GroupsGroupIDRolesRoleSetIDUsersParams{
		Limit:     &limitDefault,
		SortOrder: &sortOrderDefault,

		Context: ctx,
	}
}

// NewGETV1GroupsGroupIDRolesRoleSetIDUsersParamsWithHTTPClient creates a new GETV1GroupsGroupIDRolesRoleSetIDUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETV1GroupsGroupIDRolesRoleSetIDUsersParamsWithHTTPClient(client *http.Client) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	var (
		limitDefault     = int32(10)
		sortOrderDefault = string("Asc")
	)
	return &GETV1GroupsGroupIDRolesRoleSetIDUsersParams{
		Limit:      &limitDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GETV1GroupsGroupIDRolesRoleSetIDUsersParams contains all the parameters to send to the API endpoint
for the g e t v1 groups group Id roles role set Id users operation typically these are written to a http.Request
*/
type GETV1GroupsGroupIDRolesRoleSetIDUsersParams struct {

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
	/*RoleSetID
	  The group's role set id.

	*/
	RoleSetID int32
	/*SortOrder
	  Sorted by user group join date

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) WithTimeout(timeout time.Duration) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) WithContext(ctx context.Context) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) WithHTTPClient(client *http.Client) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) WithCursor(cursor *string) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithGroupID adds the groupID to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) WithGroupID(groupID int32) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) SetGroupID(groupID int32) {
	o.GroupID = groupID
}

// WithLimit adds the limit to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) WithLimit(limit *int32) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithRoleSetID adds the roleSetID to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) WithRoleSetID(roleSetID int32) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	o.SetRoleSetID(roleSetID)
	return o
}

// SetRoleSetID adds the roleSetId to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) SetRoleSetID(roleSetID int32) {
	o.RoleSetID = roleSetID
}

// WithSortOrder adds the sortOrder to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) WithSortOrder(sortOrder *string) *GETV1GroupsGroupIDRolesRoleSetIDUsersParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the g e t v1 groups group Id roles role set Id users params
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param roleSetId
	if err := r.SetPathParam("roleSetId", swag.FormatInt32(o.RoleSetID)); err != nil {
		return err
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