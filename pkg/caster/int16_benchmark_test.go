package caster

import "testing"

func BenchmarkToInt16(b *testing.B) {
	c := NewCaster()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ToInt16(123, -1)
		c.ToInt16(int32(123), -1)
		c.ToInt16(int64(123), -1)
		c.ToInt16(int16(12), -1)
		c.ToInt16(int8(1), -1)
		c.ToInt16(float32(123), -1)
		c.ToInt16(float64(123), -1)
		c.ToInt16("123", -1)
		c.ToInt16(make(chan int), -1)
	}
}
