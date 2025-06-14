package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOffsetAndLimit(t *testing.T) {
	env := &Env{}

	cases := []struct {
		Page     string
		PageSize string
		Offset   int
		Limit    int
	}{
		{"", "", 0, 999999},
		{"0", "20", 0, 999999},
		{"-1", "20", 0, 999999},
		{"1", "0", 0, 999999},
		{"1", "-1", 0, 999999},
		{"1", "20", 0, 20},
		{"2", "20", 20, 20},
		{"3", "20", 40, 20},
		{"4", "20", 60, 20},
		{"4", "2000", 6000, 2000},
	}
	for _, c := range cases {
		offset, limit := env.GetOffsetAndLimit(c.Page, c.PageSize)
		assert.Equal(t, c.Offset, offset)
		assert.Equal(t, c.Limit, limit)
	}
}
