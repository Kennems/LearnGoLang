package main

import "testing"

func TestTypedNilAsInterface(t *testing.T) {
	if typedNilAsInterface() == nil {
		t.Fatal("typed nil in interface should not be nil")
	}
}

func TestNilCoderValue(t *testing.T) {
	var c nilCoder
	if c != nil {
		t.Fatal("zero interface should be nil")
	}

	c = nilCoderValue()
	if c == nil {
		t.Fatal("typed nil assigned to interface should not be nil")
	}
}

func TestCompareInterfaces(t *testing.T) {
	if !compareInterfaces(any(7), any(7)) {
		t.Fatal("expected equal interfaces")
	}
}

func TestCompareMayPanic(t *testing.T) {
	_, panicked := compareMayPanic(any([]int{1}), any([]int{1}))
	if !panicked {
		t.Fatal("expected compare to panic for non-comparable dynamic type")
	}
}
