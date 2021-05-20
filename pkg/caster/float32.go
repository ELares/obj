package caster

import (
	"math"
	"strconv"
)

const (
	FLOAT64_MAXFLOAT32 = float64(math.MaxFloat32)
	FLOAT64_MINFLOAT32 = float64(math.MaxFloat32 * -1)
)

// ToInt takes an interface{} and tries to convert it to an int
// if it fails to convert the value, it will return the <def> parameter instead.
// Will attempt to cast from the following types:
// * int, int32, int64, int16, int8
// * float32, float64
// * string
func (c *Caster) ToFloat32(o interface{}, def float32) float32 {
	if val, ok := o.(int); ok {
		return float32(val)
	}

	if val, ok := o.(int32); ok {
		return float32(val)
	}

	if val, ok := o.(int64); ok {
		return float32(val)
	}

	if val, ok := o.(int16); ok {
		return float32(val)
	}

	if val, ok := o.(int8); ok {
		return float32(val)
	}

	if val, ok := o.(float32); ok {
		return val
	}

	if val, ok := o.(float64); ok {
		return c.Float64ToFloat32(val, def)
	}

	if val, ok := o.(string); ok {
		if ival, err := strconv.ParseInt(val, 10, 64); err == nil {
			return float32(ival)
		}

		if ival, err := strconv.ParseFloat(val, 64); err == nil {
			return c.Float64ToFloat32(ival, def)
		}
	}

	return def
}

// Float64ToFloat32 tries to convert float64 type to an float32
// if there's a max/min overflow then the default value
// supplied will be returned
func (c *Caster) Float64ToFloat32(f float64, def float32) float32 {
	// MAX OVERFLOW
	if f > FLOAT64_MAXFLOAT32 {
		return def
	}

	// MIN OVERFLOW
	if f < FLOAT64_MINFLOAT32 {
		return def
	}

	// MIN PRECISION OVERFLOW
	if math.Abs(f) < math.SmallestNonzeroFloat32 {
		return def
	}

	return float32(f)
}
