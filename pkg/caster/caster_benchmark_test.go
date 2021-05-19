package caster

import "testing"

func BenchmarkNewParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCaster()
	}
}

func BenchmarkToInt(b *testing.B) {
	c := NewCaster()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ToInt(123, -1)
		c.ToInt(int32(123), -1)
		c.ToInt(int64(123), -1)
		c.ToInt(int16(12), -1)
		c.ToInt(int8(1), -1)
		c.ToInt(float32(123), -1)
		c.ToInt(float64(123), -1)
		c.ToInt("123", -1)
		c.ToInt(make(chan int), -1)
	}
}
