// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GameStatus status of game
//
// swagger:model GameStatus
type GameStatus string

func NewGameStatus(value GameStatus) *GameStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GameStatus.
func (m GameStatus) Pointer() *GameStatus {
	return &m
}

const (

	// GameStatusW captures enum value "w"
	GameStatusW GameStatus = "w"

	// GameStatusG captures enum value "g"
	GameStatusG GameStatus = "g"

	// GameStatusF captures enum value "f"
	GameStatusF GameStatus = "f"
)

// for schema
var gameStatusEnum []interface{}

func init() {
	var res []GameStatus
	if err := json.Unmarshal([]byte(`["w","g","f"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gameStatusEnum = append(gameStatusEnum, v)
	}
}

func (m GameStatus) validateGameStatusEnum(path, location string, value GameStatus) error {
	if err := validate.EnumCase(path, location, value, gameStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this game status
func (m GameStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGameStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this game status based on context it is used
func (m GameStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
