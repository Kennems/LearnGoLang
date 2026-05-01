package main

import (
	"reflect"
	"testing"
)

func TestPrintArray2MutatesBackingSlice(t *testing.T) {
	data := []int{1, 2, 3}

	printArray2(data)

	if data[0] != 100 {
		t.Fatalf("data[0] = %d, want 100", data[0])
	}
}

func TestCloneSliceCreatesCopy(t *testing.T) {
	src := []int{1, 2, 3}
	got := cloneSlice(src)
	src[0] = 99

	if reflect.DeepEqual(src, got) {
		t.Fatalf("cloneSlice() returned alias: src=%v got=%v", src, got)
	}
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("cloneSlice() = %v", got)
	}
}

func TestTrimNumbers(t *testing.T) {
	got := trimNumbers([]int{0, 1, 2, 3})
	want := []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("trimNumbers() = %v, want %v", got, want)
	}
}
