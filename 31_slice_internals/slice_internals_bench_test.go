package main

import "testing"

func BenchmarkCloneSliceValues(b *testing.B) {
	src := make([]int, 1024)
	for i := range src {
		src[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cloneSliceValues(src)
	}
}
