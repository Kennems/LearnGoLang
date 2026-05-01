package main

import "testing"

func TestFieldOffset(t *testing.T) {
	if got := fieldOffset(); got == 0 {
		t.Fatal("fieldOffset() should not be zero for second field")
	}
}

func TestPointerRoundTrip(t *testing.T) {
	v := 10
	if got := *pointerRoundTrip(&v); got != 10 {
		t.Fatalf("pointerRoundTrip() = %d, want 10", got)
	}
}

func TestEscapeCandidate(t *testing.T) {
	if got := *escapeCandidate(); got != 42 {
		t.Fatalf("escapeCandidate() = %d, want 42", got)
	}
}
