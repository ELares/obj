package caster

type (
	// ICaster common methods that deals with interface{} data type
	ICaster interface {
		ToInt(o interface{}, def int) int
		ToInt32(o interface{}, def int32) int32

		IntToInt32(n int, def int32) int32
		Int64ToInt32(n int64, def int32) int32

		Float32ToInt(f float32, def int) int
		Float32ToInt32(f float32, def int32) int32

		Float64ToInt(f float64, def int) int
		Float64ToInt32(f float64, def int32) int32
	}

	// Caster implements ICaster interface
	Caster struct{}
)

// NewCast returns an implemented ICaster
func NewCaster() ICaster {
	return &Caster{}
}
