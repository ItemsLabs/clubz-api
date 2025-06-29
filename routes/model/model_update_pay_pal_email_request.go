// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelUpdatePayPalEmailRequest model update pay pal email request
//
// swagger:model model.UpdatePayPalEmailRequest
type ModelUpdatePayPalEmailRequest struct {

	// paypal email
	// Required: true
	PaypalEmail string `json:"paypal_email,omitempty"`
}

// Validate validates this model update pay pal email request
func (m *ModelUpdatePayPalEmailRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model update pay pal email request based on context it is used
func (m *ModelUpdatePayPalEmailRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelUpdatePayPalEmailRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelUpdatePayPalEmailRequest) UnmarshalBinary(b []byte) error {
	var res ModelUpdatePayPalEmailRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
