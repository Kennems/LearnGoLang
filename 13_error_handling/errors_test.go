package main

import (
	"errors"
	"testing"
)

func TestSafeSqrtNegative(t *testing.T) {
	_, err := safeSqrt(-1)
	if !errors.Is(err, ErrNegativeNumber) {
		t.Fatalf("expected ErrNegativeNumber, got %v", err)
	}
}

func TestSafeSqrtValidationError(t *testing.T) {
	_, err := safeSqrt(2)

	var validationErr *ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("expected ValidationError, got %v", err)
	}
	if validationErr.Field != "n" {
		t.Fatalf("validationErr.Field = %q, want %q", validationErr.Field, "n")
	}
}

func TestSafeSqrtSuccess(t *testing.T) {
	got, err := safeSqrt(49)
	if err != nil {
		t.Fatalf("safeSqrt() error = %v", err)
	}
	if got != 7 {
		t.Fatalf("safeSqrt() = %d, want 7", got)
	}
}
