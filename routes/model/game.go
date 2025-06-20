// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/itemslabs/clubz-api/types"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Game game
//
// swagger:model Game
type Game struct {

	// allowed actions
	AllowedActions *GameAllowedActions `json:"allowed_actions,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// match id
	MatchID string `json:"match_id,omitempty"`

	// picks
	Picks []*GamePick `json:"picks"`

	// powerups
	Powerups []*GamePowerUp `json:"powerups"`

	// premium
	Premium bool `json:"premium"`

	// rewards
	Rewards []*MatchReward `json:"rewards"`

	// Total points of the footballer pick, including powerup bonuses
	Score *types.FloatWithZero `json:"score,omitempty"`

	// status
	Status GameStatus `json:"status,omitempty"`

	// subscription tier
	SubscriptionTier SubscriptionTier `json:"subscription_tier,omitempty"`

	// swaps
	Swaps []*PlayerSwap `json:"swaps"`

	// user id
	UserID string `json:"user_id,omitempty"`

	// version of game
	Version int64 `json:"version"`
}

// Validate validates this game
func (m *Game) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePicks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionTier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwaps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Game) validateAllowedActions(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedActions) { // not required
		return nil
	}

	if m.AllowedActions != nil {
		if err := m.AllowedActions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowed_actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allowed_actions")
			}
			return err
		}
	}

	return nil
}

func (m *Game) validatePicks(formats strfmt.Registry) error {
	if swag.IsZero(m.Picks) { // not required
		return nil
	}

	for i := 0; i < len(m.Picks); i++ {
		if swag.IsZero(m.Picks[i]) { // not required
			continue
		}

		if m.Picks[i] != nil {
			if err := m.Picks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("picks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("picks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Game) validatePowerups(formats strfmt.Registry) error {
	if swag.IsZero(m.Powerups) { // not required
		return nil
	}

	for i := 0; i < len(m.Powerups); i++ {
		if swag.IsZero(m.Powerups[i]) { // not required
			continue
		}

		if m.Powerups[i] != nil {
			if err := m.Powerups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("powerups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("powerups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Game) validateRewards(formats strfmt.Registry) error {
	if swag.IsZero(m.Rewards) { // not required
		return nil
	}

	for i := 0; i < len(m.Rewards); i++ {
		if swag.IsZero(m.Rewards[i]) { // not required
			continue
		}

		if m.Rewards[i] != nil {
			if err := m.Rewards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rewards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rewards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Game) validateScore(formats strfmt.Registry) error {
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

func (m *Game) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *Game) validateSubscriptionTier(formats strfmt.Registry) error {
	if swag.IsZero(m.SubscriptionTier) { // not required
		return nil
	}

	if err := m.SubscriptionTier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscription_tier")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("subscription_tier")
		}
		return err
	}

	return nil
}

func (m *Game) validateSwaps(formats strfmt.Registry) error {
	if swag.IsZero(m.Swaps) { // not required
		return nil
	}

	for i := 0; i < len(m.Swaps); i++ {
		if swag.IsZero(m.Swaps[i]) { // not required
			continue
		}

		if m.Swaps[i] != nil {
			if err := m.Swaps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("swaps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("swaps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this game based on the context it is used
func (m *Game) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowedActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePicks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowerups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRewards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubscriptionTier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwaps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Game) contextValidateAllowedActions(ctx context.Context, formats strfmt.Registry) error {

	if m.AllowedActions != nil {

		if swag.IsZero(m.AllowedActions) { // not required
			return nil
		}

		if err := m.AllowedActions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowed_actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allowed_actions")
			}
			return err
		}
	}

	return nil
}

func (m *Game) contextValidatePicks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Picks); i++ {

		if m.Picks[i] != nil {

			if swag.IsZero(m.Picks[i]) { // not required
				return nil
			}

			if err := m.Picks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("picks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("picks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Game) contextValidatePowerups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Powerups); i++ {

		if m.Powerups[i] != nil {

			if swag.IsZero(m.Powerups[i]) { // not required
				return nil
			}

			if err := m.Powerups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("powerups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("powerups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Game) contextValidateRewards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rewards); i++ {

		if m.Rewards[i] != nil {

			if swag.IsZero(m.Rewards[i]) { // not required
				return nil
			}

			if err := m.Rewards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rewards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rewards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Game) contextValidateScore(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Game) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *Game) contextValidateSubscriptionTier(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionTier) { // not required
		return nil
	}

	if err := m.SubscriptionTier.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscription_tier")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("subscription_tier")
		}
		return err
	}

	return nil
}

func (m *Game) contextValidateSwaps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Swaps); i++ {

		if m.Swaps[i] != nil {

			if swag.IsZero(m.Swaps[i]) { // not required
				return nil
			}

			if err := m.Swaps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("swaps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("swaps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Game) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Game) UnmarshalBinary(b []byte) error {
	var res Game
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
