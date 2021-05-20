package caster

import (
	"math"
	"strconv"
)

const (
	FLOAT32_MAXINT64 = float32(math.MaxInt64)
	FLOAT64_MAXINT64 = float64(math.MaxInt64)

	FLOAT32_MININT64 = float32(math.MinInt64)
	FLOAT64_MININT64 = float64(math.MinInt64)
)

// ToInt64 takes an interface{} and tries to convert it to an int64
// if it fails to convert the value, it will return the <def> parameter instead.
// Will attempt to cast from the following types:
// * int, int32, int64, int16, int8
// * float32, float64
// * string
func (c *Caster) ToInt64(o interface{}, def int64) int64 {
	if val, ok := o.(int); ok {
		return int64(val)
	}

	if val, ok := o.(int32); ok {
		return int64(val)
	}

	if val, ok := o.(int64); ok {
		return val
	}

	if val, ok := o.(int16); ok {
		return int64(val)
	}

	if val, ok := o.(int8); ok {
		return int64(val)
	}

	if val, ok := o.(float32); ok {
		return c.Float32ToInt64(val, def)
	}

	if val, ok := o.(float64); ok {
		return c.Float64ToInt64(val, def)
	}

	if val, ok := o.(string); ok {
		if ival, err := strconv.ParseInt(val, 10, 64); err == nil {
			return ival
		}

		if ival, err := strconv.ParseFloat(val, 64); err == nil {
			return c.Float64ToInt64(ival, def)
		}
	}

	return def
}

func (c *Caster) Float32ToInt64(f float32, def int64) int64 {
	// MAX OVERFLOW
	if f > FLOAT32_MAXINT64 {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT32_MININT64 {
		return def
	}

	return int64(f)
}

func (c *Caster) Float64ToInt64(f float64, def int64) int64 {
	// MAX OVERFLOW
	if f > FLOAT64_MAXINT64 {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT64_MININT64 {
		return def
	}

	return int64(f)
}
