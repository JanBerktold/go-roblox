// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIPageResponseGroupWallPostModel Api page response group wall post model
// swagger:model ApiPageResponse[GroupWallPostModel]

type APIPageResponseGroupWallPostModel struct {

	// data
	Data APIPageResponseGroupWallPostModelData `json:"data"`

	// next page cursor
	NextPageCursor string `json:"nextPageCursor,omitempty"`

	// previous page cursor
	PreviousPageCursor string `json:"previousPageCursor,omitempty"`
}

/* polymorph ApiPageResponse[GroupWallPostModel] data false */

/* polymorph ApiPageResponse[GroupWallPostModel] nextPageCursor false */

/* polymorph ApiPageResponse[GroupWallPostModel] previousPageCursor false */

// Validate validates this Api page response group wall post model
func (m *APIPageResponseGroupWallPostModel) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *APIPageResponseGroupWallPostModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPageResponseGroupWallPostModel) UnmarshalBinary(b []byte) error {
	var res APIPageResponseGroupWallPostModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
