package apiconv

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/types"
	"github.com/go-openapi/strfmt"
	"github.com/volatiletech/null/v8"
)

func ToInt64Ptr(val int64) *int64 {
	return &val
}

func ToInt64PrtFromIntPtr(val *int) *int64 {
	if val == nil {
		return nil
	}

	var t = int64(*val)
	return &t
}

func ToStringPtr(val null.String) *string {
	if !val.Valid {
		return nil
	}

	return &val.String
}

func ToDateTime(val null.Time) *strfmt.DateTime {
	if !val.Valid {
		return nil
	}

	t := strfmt.DateTime(val.Time)
	return &t
}

func ToFloatWithZero(val float64) *types.FloatWithZero {
	t := types.FloatWithZero(val)
	return &t
}

func ToFloatWithZeroPtr(val *float64) *types.FloatWithZero {
	if val == nil {
		return nil
	}

	t := types.FloatWithZero(*val)
	return &t
}

func ToUUIDPtr(val string) *strfmt.UUID {
	uuid := strfmt.UUID(val)
	return &uuid
}
