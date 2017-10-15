// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIPageResponsePlaceModel Api page response place model
// swagger:model ApiPageResponse[PlaceModel]

type APIPageResponsePlaceModel struct {

	// data
	Data APIPageResponsePlaceModelData `json:"data"`

	// next page cursor
	NextPageCursor string `json:"nextPageCursor,omitempty"`

	// previous page cursor
	PreviousPageCursor string `json:"previousPageCursor,omitempty"`
}

/* polymorph ApiPageResponse[PlaceModel] data false */

/* polymorph ApiPageResponse[PlaceModel] nextPageCursor false */

/* polymorph ApiPageResponse[PlaceModel] previousPageCursor false */

// Validate validates this Api page response place model
func (m *APIPageResponsePlaceModel) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *APIPageResponsePlaceModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPageResponsePlaceModel) UnmarshalBinary(b []byte) error {
	var res APIPageResponsePlaceModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
