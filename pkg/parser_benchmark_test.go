package obj

import "testing"

func BenchmarkNewParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewParser()
	}
}

func BenchmarkToInt(b *testing.B) {
	ip := NewParser()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ip.ToInt(123, -1)
		ip.ToInt(int32(123), -1)
		ip.ToInt(int64(123), -1)
		ip.ToInt(int16(12), -1)
		ip.ToInt(int8(1), -1)
		ip.ToInt(float32(123), -1)
		ip.ToInt(float64(123), -1)
		ip.ToInt("123", -1)
		ip.ToInt(make(chan int), -1)
	}
}
