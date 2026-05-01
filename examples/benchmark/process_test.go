package benchmark

import (
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	data := []int{1, 2, 3, 4}

	Process(data)

	want := []int{1, 4, 9, 16}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("Process() = %v, want %v", data, want)
	}
}

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Process(data)
	}
}
