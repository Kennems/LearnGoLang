package main

import "testing"

func TestCollectElementAddresses(t *testing.T) {
	items := []item{{Value: 1}, {Value: 2}, {Value: 3}}

	got := collectElementAddresses(items)

	for i := range got {
		if got[i] != &items[i] {
			t.Fatalf("got[%d] does not point to original element", i)
		}
	}
}

func TestCollectRangeVariableAddressesAreDistinctCopies(t *testing.T) {
	items := []item{{Value: 1}, {Value: 2}, {Value: 3}}

	got := collectRangeVariableAddresses(items)

	for i := range got {
		if got[i] == &items[i] {
			t.Fatalf("got[%d] unexpectedly points to original element", i)
		}
		if got[i].Value != items[i].Value {
			t.Fatalf("got[%d].Value = %d, want %d", i, got[i].Value, items[i].Value)
		}
	}
}

func TestIncrementByRangeValueDoesNotModifyOriginal(t *testing.T) {
	items := []item{{Value: 1}, {Value: 2}}

	incrementByRangeValue(items)

	if items[0].Value != 1 || items[1].Value != 2 {
		t.Fatalf("items modified unexpectedly: %+v", items)
	}
}

func TestIncrementByIndexModifiesOriginal(t *testing.T) {
	items := []item{{Value: 1}, {Value: 2}}

	incrementByIndex(items)

	if items[0].Value != 2 || items[1].Value != 3 {
		t.Fatalf("items = %+v", items)
	}
}

func TestCaptureLoopValues(t *testing.T) {
	got := captureLoopValues([]int{1, 2, 3, 4})
	seen := make(map[int]bool, len(got))
	for _, v := range got {
		seen[v] = true
	}
	for _, want := range []int{1, 2, 3, 4} {
		if !seen[want] {
			t.Fatalf("missing captured value %d in %v", want, got)
		}
	}
}
