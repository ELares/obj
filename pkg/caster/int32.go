package caster

import (
	"math"
	"strconv"
)

const (
	INT_MAXINT32     = int(math.MaxInt32)
	INT64_MAXINT32   = int64(math.MaxInt32)
	FLOAT32_MAXINT32 = float32(math.MaxInt32)
	FLOAT64_MAXINT32 = float64(math.MaxInt32)

	INT_MININT32     = int(math.MinInt32)
	INT64_MININT32   = int64(math.MinInt32)
	FLOAT32_MININT32 = float32(math.MinInt32)
	FLOAT64_MININT32 = float64(math.MinInt32)
)

// ToInt32 takes an interface{} and tries to convert it to an int32
// if it fails to convert the value, it will return the <def> parameter instead.
// Will attempt to cast from the following types:
// * int, int32, int64, int16, int8
// * float32, float64
// * string
func (c *Caster) ToInt32(o interface{}, def int32) int32 {
	if val, ok := o.(int); ok {
		return c.IntToInt32(val, def)
	}

	if val, ok := o.(int32); ok {
		return val
	}

	if val, ok := o.(int64); ok {
		return c.Int64ToInt32(val, def)
	}

	if val, ok := o.(int16); ok {
		return int32(val)
	}

	if val, ok := o.(int8); ok {
		return int32(val)
	}

	if val, ok := o.(float32); ok {
		return c.Float32ToInt32(val, def)
	}

	if val, ok := o.(float64); ok {
		return c.Float64ToInt32(val, def)
	}

	if val, ok := o.(string); ok {
		if ival, err := strconv.ParseInt(val, 10, 64); err == nil {
			return c.Int64ToInt32(ival, def)
		}

		if ival, err := strconv.ParseFloat(val, 64); err == nil {
			return c.Float64ToInt32(ival, def)
		}
	}

	return def
}

// IntToInt32 tries to convert int type to an int32
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) IntToInt32(n int, def int32) int32 {
	// MAX OVERFLOW
	if n > INT_MAXINT32 {
		return def
	}

	// MIN OVERFLOW
	if n < INT_MININT32 {
		return def
	}

	return int32(n)
}

// Int64ToInt32 tries to convert int64 type to an int32
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) Int64ToInt32(n int64, def int32) int32 {
	// MAX OVERFLOW
	if n > INT64_MAXINT32 {
		return def
	}

	// MIN OVERFLOW
	if n < INT64_MININT32 {
		return def
	}

	return int32(n)
}

// Float32ToInt tries to convert float32 type to an int
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) Float32ToInt32(f float32, def int32) int32 {
	// MAX OVERFLOW
	if f > FLOAT32_MAXINT32 {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT32_MININT32 {
		return def
	}

	return int32(f)
}

// Float64ToInt tries to convert float64 type to an int
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) Float64ToInt32(f float64, def int32) int32 {
	// MAX OVERFLOW
	if f > FLOAT64_MAXINT32 {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT64_MININT32 {
		return def
	}

	return int32(f)
}
