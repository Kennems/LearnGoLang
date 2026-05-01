package main

import "testing"

func BenchmarkProcess(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		process(data)
	}
}

func process(data []int) {
	for i := 0; i < len(data); i++ {
		data[i] = data[i] * data[i]
	}
}
