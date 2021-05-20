package caster

import "testing"

func BenchmarkToFloat32(b *testing.B) {
	c := NewCaster()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ToFloat32(123, -1)
		c.ToFloat32(int32(123), -1)
		c.ToFloat32(int64(123), -1)
		c.ToFloat32(int16(12), -1)
		c.ToFloat32(int8(1), -1)
		c.ToFloat32(float32(123), -1)
		c.ToFloat32(float64(123), -1)
		c.ToFloat32("123", -1)
		c.ToFloat32(make(chan int), -1)
	}
}
