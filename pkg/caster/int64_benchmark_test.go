package caster

import "testing"

func BenchmarkToInt64(b *testing.B) {
	c := NewCaster()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ToInt64(123, -1)
		c.ToInt64(int32(123), -1)
		c.ToInt64(int64(123), -1)
		c.ToInt64(int16(12), -1)
		c.ToInt64(int8(1), -1)
		c.ToInt64(float32(123), -1)
		c.ToInt64(float64(123), -1)
		c.ToInt64("123", -1)
		c.ToInt64(make(chan int), -1)
	}
}
