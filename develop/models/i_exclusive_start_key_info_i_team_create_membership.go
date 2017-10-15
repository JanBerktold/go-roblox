// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IExclusiveStartKeyInfoITeamCreateMembership i exclusive start key info i team create membership
// swagger:model IExclusiveStartKeyInfo[ITeamCreateMembership]

type IExclusiveStartKeyInfoITeamCreateMembership struct {

	// count
	// Read Only: true
	Count int32 `json:"Count,omitempty"`

	// paging direction
	// Read Only: true
	PagingDirection string `json:"PagingDirection,omitempty"`

	// sort order
	// Read Only: true
	SortOrder string `json:"SortOrder,omitempty"`
}

/* polymorph IExclusiveStartKeyInfo[ITeamCreateMembership] Count false */

/* polymorph IExclusiveStartKeyInfo[ITeamCreateMembership] PagingDirection false */

/* polymorph IExclusiveStartKeyInfo[ITeamCreateMembership] SortOrder false */

// Validate validates this i exclusive start key info i team create membership
func (m *IExclusiveStartKeyInfoITeamCreateMembership) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagingDirection(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSortOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var iExclusiveStartKeyInfoITeamCreateMembershipTypePagingDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Forward","Backward"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iExclusiveStartKeyInfoITeamCreateMembershipTypePagingDirectionPropEnum = append(iExclusiveStartKeyInfoITeamCreateMembershipTypePagingDirectionPropEnum, v)
	}
}

const (
	// IExclusiveStartKeyInfoITeamCreateMembershipPagingDirectionForward captures enum value "Forward"
	IExclusiveStartKeyInfoITeamCreateMembershipPagingDirectionForward string = "Forward"
	// IExclusiveStartKeyInfoITeamCreateMembershipPagingDirectionBackward captures enum value "Backward"
	IExclusiveStartKeyInfoITeamCreateMembershipPagingDirectionBackward string = "Backward"
)

// prop value enum
func (m *IExclusiveStartKeyInfoITeamCreateMembership) validatePagingDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iExclusiveStartKeyInfoITeamCreateMembershipTypePagingDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IExclusiveStartKeyInfoITeamCreateMembership) validatePagingDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.PagingDirection) { // not required
		return nil
	}

	// value enum
	if err := m.validatePagingDirectionEnum("PagingDirection", "body", m.PagingDirection); err != nil {
		return err
	}

	return nil
}

var iExclusiveStartKeyInfoITeamCreateMembershipTypeSortOrderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Asc","Desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iExclusiveStartKeyInfoITeamCreateMembershipTypeSortOrderPropEnum = append(iExclusiveStartKeyInfoITeamCreateMembershipTypeSortOrderPropEnum, v)
	}
}

const (
	// IExclusiveStartKeyInfoITeamCreateMembershipSortOrderAsc captures enum value "Asc"
	IExclusiveStartKeyInfoITeamCreateMembershipSortOrderAsc string = "Asc"
	// IExclusiveStartKeyInfoITeamCreateMembershipSortOrderDesc captures enum value "Desc"
	IExclusiveStartKeyInfoITeamCreateMembershipSortOrderDesc string = "Desc"
)

// prop value enum
func (m *IExclusiveStartKeyInfoITeamCreateMembership) validateSortOrderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iExclusiveStartKeyInfoITeamCreateMembershipTypeSortOrderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IExclusiveStartKeyInfoITeamCreateMembership) validateSortOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.SortOrder) { // not required
		return nil
	}

	// value enum
	if err := m.validateSortOrderEnum("SortOrder", "body", m.SortOrder); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IExclusiveStartKeyInfoITeamCreateMembership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IExclusiveStartKeyInfoITeamCreateMembership) UnmarshalBinary(b []byte) error {
	var res IExclusiveStartKeyInfoITeamCreateMembership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
