package caster

type (
	// ICaster common methods that deals with interface{} data type
	ICaster interface {
		ToInt(o interface{}, def int) int
		ToInt8(o interface{}, def int8) int8
		ToInt16(o interface{}, def int16) int16
		ToInt32(o interface{}, def int32) int32
		ToInt64(o interface{}, def int64) int64
		ToFloat32(o interface{}, def float32) float32

		IntToInt8(n int, def int8) int8
		IntToInt16(n int, def int16) int16
		IntToInt32(n int, def int32) int32

		Int16ToInt8(n int16, def int8) int8

		Int32ToInt8(n int32, def int8) int8
		Int32ToInt16(n int32, def int16) int16

		Int64ToInt8(n int64, def int8) int8
		Int64ToInt16(n int64, def int16) int16
		Int64ToInt32(n int64, def int32) int32

		Float32ToInt(f float32, def int) int
		Float32ToInt8(f float32, def int8) int8
		Float32ToInt16(f float32, def int16) int16
		Float32ToInt32(f float32, def int32) int32
		Float32ToInt64(f float32, def int64) int64

		Float64ToInt(f float64, def int) int
		Float64ToInt8(f float64, def int8) int8
		Float64ToInt16(f float64, def int16) int16
		Float64ToInt32(f float64, def int32) int32
		Float64ToInt64(f float64, def int64) int64
		Float64ToFloat32(f float64, def float32) float32
	}

	// Caster implements ICaster interface
	Caster struct{}
)

// NewCast returns an implemented ICaster
func NewCaster() ICaster {
	return &Caster{}
}
