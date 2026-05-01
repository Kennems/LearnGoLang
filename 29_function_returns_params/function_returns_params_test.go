package main

import (
	"reflect"
	"testing"
)

func TestAddOneByValue(t *testing.T) {
	n := 1
	addOneByValue(n)
	if n != 1 {
		t.Fatalf("n = %d, want 1", n)
	}
}

func TestMutateSliceElements(t *testing.T) {
	values := []int{1, 2, 3}
	mutateSliceElements(values)
	if !reflect.DeepEqual(values, []int{2, 3, 4}) {
		t.Fatalf("values = %v", values)
	}
}

func TestAppendValue(t *testing.T) {
	values := []int{1, 2}
	newValues := appendValue(values, 3)

	if !reflect.DeepEqual(values, []int{1, 2}) {
		t.Fatalf("original values unexpectedly changed: %v", values)
	}
	if !reflect.DeepEqual(newValues, []int{1, 2, 3}) {
		t.Fatalf("newValues = %v", newValues)
	}
}

func TestAppendValueByPointer(t *testing.T) {
	values := []int{1, 2}
	appendValueByPointer(&values, 3)
	if !reflect.DeepEqual(values, []int{1, 2, 3}) {
		t.Fatalf("values = %v", values)
	}
}

func TestSplitAndSum(t *testing.T) {
	a, b, sum := splitAndSum(2, 3)
	if a != 2 || b != 3 || sum != 5 {
		t.Fatalf("got %d %d %d", a, b, sum)
	}
}

func TestDeferredReturns(t *testing.T) {
	if got := deferredNamedReturn(); got != 2 {
		t.Fatalf("deferredNamedReturn() = %d, want 2", got)
	}
	if got := deferredUnnamedReturn(); got != 1 {
		t.Fatalf("deferredUnnamedReturn() = %d, want 1", got)
	}
}
