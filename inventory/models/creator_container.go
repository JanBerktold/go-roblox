// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreatorContainer creator container
// swagger:model CreatorContainer

type CreatorContainer struct {

	// creator Id
	CreatorID int64 `json:"CreatorId,omitempty"`

	// creator type
	CreatorType string `json:"CreatorType,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

/* polymorph CreatorContainer CreatorId false */

/* polymorph CreatorContainer CreatorType false */

/* polymorph CreatorContainer Name false */

// Validate validates this creator container
func (m *CreatorContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CreatorContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatorContainer) UnmarshalBinary(b []byte) error {
	var res CreatorContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}