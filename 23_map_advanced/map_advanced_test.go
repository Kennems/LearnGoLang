package main

import (
	"reflect"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	got := sortedKeys(map[int]string{3: "c", 1: "a", 2: "b"})
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("sortedKeys() = %v, want %v", got, want)
	}
}

func TestDeleteEvens(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	deleteEvens(m)
	if len(m) != 2 {
		t.Fatalf("len(m) = %d, want 2", len(m))
	}
}
