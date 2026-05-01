package main

import "testing"

func TestAssertString(t *testing.T) {
	if got, ok := assertString("go"); !ok || got != "go" {
		t.Fatalf("assertString() = %q, %v", got, ok)
	}
}

func TestConvertInt64(t *testing.T) {
	if got := convertInt64(7); got != 7 {
		t.Fatalf("convertInt64() = %d, want 7", got)
	}
}

func TestCompareInterfaces(t *testing.T) {
	if !compareInterfaces(any(1), any(1)) {
		t.Fatal("expected interface values to be equal")
	}
}

func TestStringerImplementation(t *testing.T) {
	var s stringer = userName("ken")
	if s.String() != "ken" {
		t.Fatalf("s.String() = %q", s.String())
	}
}
