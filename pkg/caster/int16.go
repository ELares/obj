package caster

import (
	"math"
	"strconv"
)

const (
	INT_MAXINT16     = int(math.MaxInt16)
	INT32_MAXINT16   = int32(math.MaxInt16)
	INT64_MAXINT16   = int64(math.MaxInt16)
	FLOAT32_MAXINT16 = float32(math.MaxInt16)
	FLOAT64_MAXINT16 = float64(math.MaxInt16)

	INT_MININT16     = int(math.MinInt16)
	INT32_MININT16   = int32(math.MinInt16)
	INT64_MININT16   = int64(math.MinInt16)
	FLOAT32_MININT16 = float32(math.MinInt16)
	FLOAT64_MININT16 = float64(math.MinInt16)
)

// ToInt takes an interface{} and tries to convert it to an int16
// if it fails to convert the value, it will return the <def> parameter instead.
// Will attempt to cast from the following types:
// * int, int32, int64, int16, int8
// * float32, float64
// * string
func (c *Caster) ToInt16(o interface{}, def int16) int16 {
	if val, ok := o.(int); ok {
		return c.IntToInt16(val, def)
	}

	if val, ok := o.(int32); ok {
		return c.Int32ToInt16(val, def)
	}

	if val, ok := o.(int64); ok {
		return c.Int64ToInt16(val, def)
	}

	if val, ok := o.(int16); ok {
		return val
	}

	if val, ok := o.(int8); ok {
		return int16(val)
	}

	if val, ok := o.(float32); ok {
		return c.Float32ToInt16(val, def)
	}

	if val, ok := o.(float64); ok {
		return c.Float64ToInt16(val, def)
	}

	if val, ok := o.(string); ok {
		if ival, err := strconv.ParseInt(val, 10, 64); err == nil {
			return c.Int64ToInt16(ival, def)
		}

		if ival, err := strconv.ParseFloat(val, 64); err == nil {
			return c.Float64ToInt16(ival, def)
		}
	}

	return def
}

func (c *Caster) IntToInt16(n int, def int16) int16 {
	// MAX OVERFLOW
	if n > INT_MAXINT16 {
		return def
	}

	// MIN OVERFLOW
	if n < INT_MININT16 {
		return def
	}

	return int16(n)
}

func (c *Caster) Int32ToInt16(n int32, def int16) int16 {
	// MAX OVERFLOW
	if n > INT32_MAXINT16 {
		return def
	}

	// MIN OVERFLOW
	if n < INT32_MININT16 {
		return def
	}

	return int16(n)
}

func (c *Caster) Int64ToInt16(n int64, def int16) int16 {
	// MAX OVERFLOW
	if n > INT64_MAXINT16 {
		return def
	}

	// MIN OVERFLOW
	if n < INT64_MININT16 {
		return def
	}

	return int16(n)
}

// Float32ToInt8 tries to convert float32 type to an int8
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) Float32ToInt16(f float32, def int16) int16 {
	// MAX OVERFLOW
	if f > FLOAT32_MAXINT16 {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT32_MININT16 {
		return def
	}

	return int16(f)
}

// Float64ToInt8 tries to convert float64 type to an int8
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) Float64ToInt16(f float64, def int16) int16 {
	// MAX OVERFLOW
	if f > FLOAT64_MAXINT16 {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT64_MININT16 {
		return def
	}

	return int16(f)
}
