package caster

import (
	"math"
	"strconv"
)

const (
	INT_MAXINT8     = int(math.MaxInt8)
	INT16_MAXINT8   = int16(math.MaxInt8)
	INT32_MAXINT8   = int32(math.MaxInt8)
	INT64_MAXINT8   = int64(math.MaxInt8)
	FLOAT32_MAXINT8 = float32(math.MaxInt8)
	FLOAT64_MAXINT8 = float64(math.MaxInt8)

	INT_MININT8     = int(math.MinInt8)
	INT16_MININT8   = int16(math.MinInt8)
	INT32_MININT8   = int32(math.MinInt8)
	INT64_MININT8   = int64(math.MinInt8)
	FLOAT32_MININT8 = float32(math.MinInt8)
	FLOAT64_MININT8 = float64(math.MinInt8)
)

// ToInt takes an interface{} and tries to convert it to an int8
// if it fails to convert the value, it will return the <def> parameter instead.
// Will attempt to cast from the following types:
// * int, int32, int64, int16, int8
// * float32, float64
// * string
func (c *Caster) ToInt8(o interface{}, def int8) int8 {
	if val, ok := o.(int); ok {
		return c.IntToInt8(val, def)
	}

	if val, ok := o.(int32); ok {
		return c.Int32ToInt8(val, def)
	}

	if val, ok := o.(int64); ok {
		return c.Int64ToInt8(val, def)
	}

	if val, ok := o.(int16); ok {
		return c.Int16ToInt8(val, def)
	}

	if val, ok := o.(int8); ok {
		return val
	}

	if val, ok := o.(float32); ok {
		return c.Float32ToInt8(val, def)
	}

	if val, ok := o.(float64); ok {
		return c.Float64ToInt8(val, def)
	}

	if val, ok := o.(string); ok {
		if ival, err := strconv.ParseInt(val, 10, 64); err == nil {
			return c.Int64ToInt8(ival, def)
		}

		if ival, err := strconv.ParseFloat(val, 64); err == nil {
			return c.Float64ToInt8(ival, def)
		}
	}

	return def
}

func (c *Caster) IntToInt8(n int, def int8) int8 {
	// MAX OVERFLOW
	if n > INT_MAXINT8 {
		return def
	}

	// MIN OVERFLOW
	if n < INT_MININT8 {
		return def
	}

	return int8(n)
}

func (c *Caster) Int16ToInt8(n int16, def int8) int8 {
	// MAX OVERFLOW
	if n > INT16_MAXINT8 {
		return def
	}

	// MIN OVERFLOW
	if n < INT16_MININT8 {
		return def
	}

	return int8(n)
}

func (c *Caster) Int32ToInt8(n int32, def int8) int8 {
	// MAX OVERFLOW
	if n > INT32_MAXINT8 {
		return def
	}

	// MIN OVERFLOW
	if n < INT32_MININT8 {
		return def
	}

	return int8(n)
}

func (c *Caster) Int64ToInt8(n int64, def int8) int8 {
	// MAX OVERFLOW
	if n > INT64_MAXINT8 {
		return def
	}

	// MIN OVERFLOW
	if n < INT64_MININT8 {
		return def
	}

	return int8(n)
}

// Float32ToInt8 tries to convert float32 type to an int8
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) Float32ToInt8(f float32, def int8) int8 {
	// MAX OVERFLOW
	if f > FLOAT32_MAXINT8 {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT32_MININT8 {
		return def
	}

	return int8(f)
}

// Float64ToInt8 tries to convert float64 type to an int8
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) Float64ToInt8(f float64, def int8) int8 {
	// MAX OVERFLOW
	if f > FLOAT64_MAXINT8 {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT64_MININT8 {
		return def
	}

	return int8(f)
}
