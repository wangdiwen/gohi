package gohi

import (
	"testing"
)

func BenchmarkInc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Inc(100)
	}
	b.StopTimer()
}
