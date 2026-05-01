package main

import (
	"reflect"
	"testing"
)

func TestAlternatingSequence(t *testing.T) {
	got := alternatingSequence(6)
	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("alternatingSequence() = %v, want %v", got, want)
	}
}

func TestWorkerPoolCount(t *testing.T) {
	if got := workerPoolCount(100, 10); got != 100 {
		t.Fatalf("workerPoolCount() = %d, want 100", got)
	}
}
