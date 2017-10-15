// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PlaceModel A model containing information about a place
// swagger:model PlaceModel

type PlaceModel struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// universe Id
	UniverseID int64 `json:"universeId,omitempty"`
}

/* polymorph PlaceModel description false */

/* polymorph PlaceModel id false */

/* polymorph PlaceModel name false */

/* polymorph PlaceModel universeId false */

// Validate validates this place model
func (m *PlaceModel) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PlaceModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlaceModel) UnmarshalBinary(b []byte) error {
	var res PlaceModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}