package caster

// ToInt64 takes an interface{} and tries to convert it to an int64
// if it fails to convert the value, it will return the <def> parameter instead.
// Will attempt to cast from the following types:
// * int, int32, int64, int16, int8
// * float32, float64
// * string
func (c *Caster) ToInt64(o interface{}, def int64) int64 {
	return int64(c.ToInt(o, int(def)))
}
