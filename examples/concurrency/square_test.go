package concurrency

import "testing"

func TestGetSquare(t *testing.T) {
	if got := GetSquare(12); got != 144 {
		t.Fatalf("GetSquare(12) = %d, want 144", got)
	}
}
