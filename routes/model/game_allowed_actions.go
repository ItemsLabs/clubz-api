// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GameAllowedActions game allowed actions
//
// swagger:model GameAllowedActions
type GameAllowedActions struct {

	// bonus powerup left
	BonusPowerupLeft int64 `json:"bonus_powerup_left"`

	// powerup left
	PowerupLeft []int64 `json:"powerup_left"`

	// swaps left
	SwapsLeft []int64 `json:"swaps_left"`
}

// Validate validates this game allowed actions
func (m *GameAllowedActions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this game allowed actions based on context it is used
func (m *GameAllowedActions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GameAllowedActions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GameAllowedActions) UnmarshalBinary(b []byte) error {
	var res GameAllowedActions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
