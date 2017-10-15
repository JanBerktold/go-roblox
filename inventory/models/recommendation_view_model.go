// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RecommendationViewModel recommendation view model
// swagger:model RecommendationViewModel

type RecommendationViewModel struct {

	// creator
	Creator *CreatorContainer `json:"Creator,omitempty"`

	// item
	Item *ItemContainer `json:"Item,omitempty"`

	// product
	Product *ProductContainer `json:"Product,omitempty"`

	// thumbnail
	Thumbnail *ThumbResultContainer `json:"Thumbnail,omitempty"`
}

/* polymorph RecommendationViewModel Creator false */

/* polymorph RecommendationViewModel Item false */

/* polymorph RecommendationViewModel Product false */

/* polymorph RecommendationViewModel Thumbnail false */

// Validate validates this recommendation view model
func (m *RecommendationViewModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreator(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateItem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThumbnail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecommendationViewModel) validateCreator(formats strfmt.Registry) error {

	if swag.IsZero(m.Creator) { // not required
		return nil
	}

	if m.Creator != nil {

		if err := m.Creator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Creator")
			}
			return err
		}
	}

	return nil
}

func (m *RecommendationViewModel) validateItem(formats strfmt.Registry) error {

	if swag.IsZero(m.Item) { // not required
		return nil
	}

	if m.Item != nil {

		if err := m.Item.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Item")
			}
			return err
		}
	}

	return nil
}

func (m *RecommendationViewModel) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {

		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Product")
			}
			return err
		}
	}

	return nil
}

func (m *RecommendationViewModel) validateThumbnail(formats strfmt.Registry) error {

	if swag.IsZero(m.Thumbnail) { // not required
		return nil
	}

	if m.Thumbnail != nil {

		if err := m.Thumbnail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Thumbnail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecommendationViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecommendationViewModel) UnmarshalBinary(b []byte) error {
	var res RecommendationViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
