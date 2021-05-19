package caster

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasterToInt8(t *testing.T) {
	testcases := []struct {
		name string

		input    interface{}
		inputDef int8

		exp int8
	}{
		{
			name:     "string_positive",
			input:    "10",
			inputDef: -1,
			exp:      10,
		},

		{
			name:     "string_negative",
			input:    "-10",
			inputDef: -1,
			exp:      -10,
		},

		{
			name:     "string_float",
			input:    "-10.57485464987463564654654",
			inputDef: -1,
			exp:      -10,
		},

		{
			name:     "string_max_int_64",
			input:    strconv.FormatInt(math.MaxInt64, 10),
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "string_max_int_32",
			input:    strconv.FormatInt(math.MaxInt32, 10),
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "string_max_float_32",
			input:    strconv.FormatFloat(math.MaxFloat32, 'f', 32, 32),
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "string_max_float_64",
			input:    strconv.FormatFloat(math.MaxFloat64, 'f', 64, 64),
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "string_min_float_32",
			input:    strconv.FormatFloat(math.MaxFloat32*-1, 'f', 32, 32),
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "string_min_float_64",
			input:    strconv.FormatFloat(math.MaxFloat64*-1, 'f', 64, 64),
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "float64_max",
			input:    math.MaxFloat64,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "float64_min",
			input:    math.MaxFloat64 * -1,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "float32_max",
			input:    math.MaxFloat32,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "float32_min",
			input:    math.MaxFloat32 * -1,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "int32_max",
			input:    math.MaxInt32,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "int32_min",
			input:    math.MinInt32,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "int64_max",
			input:    math.MaxInt64,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "int64_min",
			input:    math.MinInt64,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "int16_max",
			input:    math.MaxInt16,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "int16_min",
			input:    math.MinInt16,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "int8_max",
			input:    math.MaxInt8,
			inputDef: -1,
			exp:      math.MaxInt8,
		},

		{
			name:     "int8_min",
			input:    math.MinInt8,
			inputDef: -1,
			exp:      math.MinInt8,
		},

		{
			name:     "nil",
			input:    nil,
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "channel",
			input:    make(chan int),
			inputDef: -1,
			exp:      -1,
		},

		{
			name:     "empty_struct",
			input:    struct{}{},
			inputDef: -1,
			exp:      -1,
		},
	}

	c := NewCaster()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := c.ToInt8(tc.input, tc.inputDef)
			assert.Equal(t, tc.exp, result)
		})
	}
}
