// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelPlayerSwap model player swap
//
// swagger:model model.PlayerSwap
type ModelPlayerSwap struct {

	// executed
	// Format: date-time
	Executed string `json:"executed,omitempty"`

	// game id
	GameID string `json:"game_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// player id in
	PlayerIDIn string `json:"player_id_in,omitempty"`

	// player id out
	PlayerIDOut string `json:"player_id_out,omitempty"`
}

// Validate validates this model player swap
func (m *ModelPlayerSwap) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model player swap based on context it is used
func (m *ModelPlayerSwap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelPlayerSwap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelPlayerSwap) UnmarshalBinary(b []byte) error {
	var res ModelPlayerSwap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
