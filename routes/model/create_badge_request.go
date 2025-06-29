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

// CreateBadgeRequest create badge request
//
// swagger:model CreateBadgeRequest
type CreateBadgeRequest struct {

	// description
	// Example: Awarded to top players.
	// Required: true
	Description *string `json:"description"`

	// image
	// Example: https://example.com/image.png
	// Required: true
	Image *string `json:"image"`

	// name
	// Example: Champion Badge
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
	// Example: gold
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this create badge request
func (m *CreateBadgeRequest) Validate(formats strfmt.Registry) error {
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

func (m *CreateBadgeRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *CreateBadgeRequest) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *CreateBadgeRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateBadgeRequest) validatePoints(formats strfmt.Registry) error {

	if err := validate.Required("points", "body", m.Points); err != nil {
		return err
	}

	return nil
}

func (m *CreateBadgeRequest) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CreateBadgeRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create badge request based on context it is used
func (m *CreateBadgeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateBadgeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBadgeRequest) UnmarshalBinary(b []byte) error {
	var res CreateBadgeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
