package goTest1

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		Add(1, 2)
	}
	b.ReportMetric(1, "ops")
	b.ReportAllocs()
}
