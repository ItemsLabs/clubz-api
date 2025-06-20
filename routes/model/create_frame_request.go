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

// CreateFrameRequest create frame request
//
// swagger:model CreateFrameRequest
type CreateFrameRequest struct {

	// description
	// Example: Special Silver Frame.
	// Required: true
	Description *string `json:"description"`

	// image
	// Example: https://example.com/image.png
	// Required: true
	Image *string `json:"image"`

	// name
	// Example: Silver Frame
	// Required: true
	Name *string `json:"name"`

	// points
	// Example: 100
	// Required: true
	Points *int64 `json:"points"`

	// status
	// Example: active
	// Required: true
	Status *string `json:"status"`

	// type
	// Example: decorative
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this create frame request
func (m *CreateFrameRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFrameRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *CreateFrameRequest) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *CreateFrameRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateFrameRequest) validatePoints(formats strfmt.Registry) error {

	if err := validate.Required("points", "body", m.Points); err != nil {
		return err
	}

	return nil
}

func (m *CreateFrameRequest) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CreateFrameRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create frame request based on context it is used
func (m *CreateFrameRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateFrameRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFrameRequest) UnmarshalBinary(b []byte) error {
	var res CreateFrameRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
