package caster

import "testing"

func BenchmarkToInt32(b *testing.B) {
	c := NewCaster()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ToInt32(123, -1)
		c.ToInt32(int32(123), -1)
		c.ToInt32(int64(123), -1)
		c.ToInt32(int16(12), -1)
		c.ToInt32(int8(1), -1)
		c.ToInt32(float32(123), -1)
		c.ToInt32(float64(123), -1)
		c.ToInt32("123", -1)
		c.ToInt32(make(chan int), -1)
	}
}
