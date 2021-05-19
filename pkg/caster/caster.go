package caster

type (
	// ICaster common methods that deals with interface{} data type
	ICaster interface {
		ToInt(o interface{}, def int) int
		ToInt8(o interface{}, def int8) int8
		ToInt32(o interface{}, def int32) int32
		ToInt64(o interface{}, def int64) int64

		IntToInt8(n int, def int8) int8
		IntToInt32(n int, def int32) int32

		Int16ToInt8(n int16, def int8) int8
		Int32ToInt8(n int32, def int8) int8

		Int64ToInt8(n int64, def int8) int8
		Int64ToInt32(n int64, def int32) int32

		Float32ToInt(f float32, def int) int
		Float32ToInt8(f float32, def int8) int8
		Float32ToInt32(f float32, def int32) int32

		Float64ToInt(f float64, def int) int
		Float64ToInt8(f float64, def int8) int8
		Float64ToInt32(f float64, def int32) int32
	}

	// Caster implements ICaster interface
	Caster struct{}
)

// NewCast returns an implemented ICaster
func NewCaster() ICaster {
	return &Caster{}
}
