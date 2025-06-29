// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelLiveGamePlayerStat model live game player stat
//
// swagger:model model.LiveGamePlayerStat
type ModelLiveGamePlayerStat struct {

	// action summary
	ActionSummary []*ModelActionSummary `json:"action_summary"`

	// last 10 min points
	Last10MinPoints float64 `json:"last_10_min_points,omitempty"`

	// point distribution
	PointDistribution []*ModelPointBucket `json:"point_distribution"`

	// selected percent
	SelectedPercent float64 `json:"selected_percent,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this model live game player stat
func (m *ModelLiveGamePlayerStat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePointDistribution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelLiveGamePlayerStat) validateActionSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionSummary) { // not required
		return nil
	}

	for i := 0; i < len(m.ActionSummary); i++ {
		if swag.IsZero(m.ActionSummary[i]) { // not required
			continue
		}

		if m.ActionSummary[i] != nil {
			if err := m.ActionSummary[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_summary" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_summary" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelLiveGamePlayerStat) validatePointDistribution(formats strfmt.Registry) error {
	if swag.IsZero(m.PointDistribution) { // not required
		return nil
	}

	for i := 0; i < len(m.PointDistribution); i++ {
		if swag.IsZero(m.PointDistribution[i]) { // not required
			continue
		}

		if m.PointDistribution[i] != nil {
			if err := m.PointDistribution[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("point_distribution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("point_distribution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this model live game player stat based on the context it is used
func (m *ModelLiveGamePlayerStat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePointDistribution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelLiveGamePlayerStat) contextValidateActionSummary(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ActionSummary); i++ {

		if m.ActionSummary[i] != nil {

			if swag.IsZero(m.ActionSummary[i]) { // not required
				return nil
			}

			if err := m.ActionSummary[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_summary" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_summary" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelLiveGamePlayerStat) contextValidatePointDistribution(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PointDistribution); i++ {

		if m.PointDistribution[i] != nil {

			if swag.IsZero(m.PointDistribution[i]) { // not required
				return nil
			}

			if err := m.PointDistribution[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("point_distribution" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("point_distribution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelLiveGamePlayerStat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelLiveGamePlayerStat) UnmarshalBinary(b []byte) error {
	var res ModelLiveGamePlayerStat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
