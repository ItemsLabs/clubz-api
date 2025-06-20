// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/itemslabs/clubz-api/types"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Action action that we score
//
// swagger:model Action
type Action struct {

	// category
	Category string `json:"category"`

	// description
	Description string `json:"description"`

	// icon
	Icon string `json:"icon"`

	// id
	ID int64 `json:"id"`

	// name
	Name string `json:"name"`

	// score
	Score *types.FloatWithZero `json:"score,omitempty"`
}

// Validate validates this action
func (m *Action) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Action) validateScore(formats strfmt.Registry) error {
	if swag.IsZero(m.Score) { // not required
		return nil
	}

	if m.Score != nil {
		if err := m.Score.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("score")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("score")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this action based on the context it is used
func (m *Action) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Action) contextValidateScore(ctx context.Context, formats strfmt.Registry) error {

	if m.Score != nil {

		if swag.IsZero(m.Score) { // not required
			return nil
		}

		if err := m.Score.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("score")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("score")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Action) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Action) UnmarshalBinary(b []byte) error {
	var res Action
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
