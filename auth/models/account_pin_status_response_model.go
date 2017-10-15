// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountPinStatusResponseModel A class representing the status of Account Pin.
// swagger:model AccountPinStatusResponseModel

type AccountPinStatusResponseModel struct {

	// Gets or sets a value indicating whether this account pin is enabled.
	IsEnabled bool `json:"isEnabled,omitempty"`

	// Returns the number of seconds left the account pin is unlocked until.
	UnlockedUntil float64 `json:"unlockedUntil,omitempty"`
}

/* polymorph AccountPinStatusResponseModel isEnabled false */

/* polymorph AccountPinStatusResponseModel unlockedUntil false */

// Validate validates this account pin status response model
func (m *AccountPinStatusResponseModel) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AccountPinStatusResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPinStatusResponseModel) UnmarshalBinary(b []byte) error {
	var res AccountPinStatusResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}