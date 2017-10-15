// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APILegacyPageResponseRecommendationViewModel Api legacy page response recommendation view model
// swagger:model ApiLegacyPageResponse[RecommendationViewModel]

type APILegacyPageResponseRecommendationViewModel struct {

	// data
	Data APILegacyPageResponseRecommendationViewModelData `json:"data"`

	// total
	Total int64 `json:"total,omitempty"`
}

/* polymorph ApiLegacyPageResponse[RecommendationViewModel] data false */

/* polymorph ApiLegacyPageResponse[RecommendationViewModel] total false */

// Validate validates this Api legacy page response recommendation view model
func (m *APILegacyPageResponseRecommendationViewModel) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *APILegacyPageResponseRecommendationViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APILegacyPageResponseRecommendationViewModel) UnmarshalBinary(b []byte) error {
	var res APILegacyPageResponseRecommendationViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}