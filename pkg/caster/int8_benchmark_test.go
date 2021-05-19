package caster

import "testing"

func BenchmarkToInt8(b *testing.B) {
	c := NewCaster()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ToInt8(123, -1)
		c.ToInt8(int32(123), -1)
		c.ToInt8(int64(123), -1)
		c.ToInt8(int16(12), -1)
		c.ToInt8(int8(1), -1)
		c.ToInt8(float32(123), -1)
		c.ToInt8(float64(123), -1)
		c.ToInt8("123", -1)
		c.ToInt8(make(chan int), -1)
	}
}
