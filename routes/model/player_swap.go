// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlayerSwap player swap
//
// swagger:model PlayerSwap
type PlayerSwap struct {

	// executed
	// Format: date-time
	Executed strfmt.DateTime `json:"executed,omitempty"`

	// game id
	GameID string `json:"game_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// player id in
	PlayerIDIn string `json:"player_id_in,omitempty"`

	// player id out
	PlayerIDOut string `json:"player_id_out,omitempty"`
}

// Validate validates this player swap
func (m *PlayerSwap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecuted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerSwap) validateExecuted(formats strfmt.Registry) error {
	if swag.IsZero(m.Executed) { // not required
		return nil
	}

	if err := validate.FormatOf("executed", "body", "date-time", m.Executed.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this player swap based on context it is used
func (m *PlayerSwap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlayerSwap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerSwap) UnmarshalBinary(b []byte) error {
	var res PlayerSwap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
