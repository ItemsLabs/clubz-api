// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NullFloat64 null float64
//
// swagger:model null.Float64
type NullFloat64 struct {

	// float64
	Float64 float64 `json:"float64,omitempty"`

	// valid
	Valid bool `json:"valid,omitempty"`
}

// Validate validates this null float64
func (m *NullFloat64) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this null float64 based on context it is used
func (m *NullFloat64) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NullFloat64) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NullFloat64) UnmarshalBinary(b []byte) error {
	var res NullFloat64
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
