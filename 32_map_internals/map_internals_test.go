package main

import (
	"reflect"
	"testing"
)

func TestSortedMapKeys(t *testing.T) {
	got := sortedMapKeys(buildIntMap())
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("sortedMapKeys() = %v, want %v", got, want)
	}
}

func TestDeleteDuringRange(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	deleteDuringRange(m)

	if len(m) != 2 {
		t.Fatalf("len(m) = %d, want 2", len(m))
	}
	if _, ok := m[2]; ok {
		t.Fatal("expected key 2 to be deleted")
	}
	if _, ok := m[4]; ok {
		t.Fatal("expected key 4 to be deleted")
	}
}

func TestMapLookup(t *testing.T) {
	v, ok := mapLookup(map[string]int{"go": 1}, "go")
	if !ok || v != 1 {
		t.Fatalf("mapLookup() = %d, %v", v, ok)
	}
}

func TestMapLengthAfterDelete(t *testing.T) {
	if got := mapLengthAfterDelete(); got != 2 {
		t.Fatalf("mapLengthAfterDelete() = %d, want 2", got)
	}
}
