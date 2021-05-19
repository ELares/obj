package caster

import "testing"

func BenchmarkNewCaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCaster()
	}
}
