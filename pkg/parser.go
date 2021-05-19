package obj

import (
	"math"
	"strconv"
)

type (
	IParser interface {
		ToInt(o interface{}, def int) int
	}

	// Parse implements IParse interface
	Parser struct{}
)

func NewParser() IParser {
	return &Parser{}
}

// ToInt takes an in// * int64erface{} and tries to convert it to an int
// if it fails to convert the value, it will return the <def> parameter instead
// Will attempt to cast from the following types:
// * int, int32, int64, int16, int8
// * float32, float64
// * string
func (p *Parser) ToInt(o interface{}, def int) int {
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
		return p.Float32ToInt(val, def)
	}

	if val, ok := o.(float64); ok {
		return p.Float64ToInt(val, def)
	}

	if val, ok := o.(string); ok {
		if ival, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int(ival)
		}

		if ival, err := strconv.ParseFloat(val, 64); err == nil {
			return p.Float64ToInt(ival, def)
		}
	}

	return def
}

func (p *Parser) Float32ToInt(f float32, def int) int {
	// MAX OVERFLOW
	if f > float32(math.MaxInt64) {
		return def
	}

	// MIN OVERFLOW
	if f < float32(math.MinInt64) {
		return def
	}

	return int(f)
}

func (p *Parser) Float64ToInt(f float64, def int) int {
	// MAX OVERFLOW
	if f > float64(math.MaxInt64) {
		return def
	}

	// MIN OVERFLOW
	if f < float64(math.MinInt64) {
		return def
	}

	return int(f)
}
