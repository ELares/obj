package caster

import (
	"math"
	"strconv"
)

const (
	FLOAT32_MAXINT = float32(math.MaxInt64)
	FLOAT64_MAXINT = float64(math.MaxInt64)

	FLOAT32_MININT = float32(math.MinInt64)
	FLOAT64_MININT = float64(math.MinInt64)
)

// ToInt takes an interface{} and tries to convert it to an int
// if it fails to convert the value, it will return the <def> parameter instead.
// Will attempt to cast from the following types:
// * int, int32, int64, int16, int8
// * float32, float64
// * string
func (c *Caster) ToInt(o interface{}, def int) int {
	if val, ok := o.(int); ok {
		return val
	}

	if val, ok := o.(int32); ok {
		return int(val)
	}

	if val, ok := o.(int64); ok {
		return int(val)
	}

	if val, ok := o.(int16); ok {
		return int(val)
	}

	if val, ok := o.(int8); ok {
		return int(val)
	}

	if val, ok := o.(float32); ok {
		return c.Float32ToInt(val, def)
	}

	if val, ok := o.(float64); ok {
		return c.Float64ToInt(val, def)
	}

	if val, ok := o.(string); ok {
		if ival, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int(ival)
		}

		if ival, err := strconv.ParseFloat(val, 64); err == nil {
			return c.Float64ToInt(ival, def)
		}
	}

	return def
}

// Float32ToInt tries to convert float32 type to an int
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) Float32ToInt(f float32, def int) int {
	// MAX OVERFLOW
	if f > FLOAT32_MAXINT {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT32_MININT {
		return def
	}

	return int(f)
}

// Float64ToInt tries to convert float64 type to an int
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) Float64ToInt(f float64, def int) int {
	// MAX OVERFLOW
	if f > FLOAT64_MAXINT {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT64_MININT {
		return def
	}

	return int(f)
}
