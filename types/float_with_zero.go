package types

import (
	"context"
	"github.com/go-openapi/strfmt"
	"strconv"
)

type FloatWithZero float64

func (f FloatWithZero) MarshalJSON() ([]byte, error) {
	if float64(f) == float64(int(f)) {
		return []byte(strconv.FormatFloat(float64(f), 'f', 1, 32)), nil
	}
	return []byte(strconv.FormatFloat(float64(f), 'f', -1, 32)), nil
}

func (f FloatWithZero) Validate(formats strfmt.Registry) error {
	return nil
}

func (f FloatWithZero) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
