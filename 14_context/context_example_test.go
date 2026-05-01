package main

import (
	"context"
	"testing"
	"time"
)

func TestRunWithTimeoutExceeded(t *testing.T) {
	err := runWithTimeout(context.Background(), 10*time.Millisecond, 30*time.Millisecond)
	if !isDeadlineExceeded(err) {
		t.Fatalf("expected deadline exceeded, got %v", err)
	}
}

func TestRunWithTimeoutSuccess(t *testing.T) {
	err := runWithTimeout(context.Background(), 30*time.Millisecond, 5*time.Millisecond)
	if err != nil {
		t.Fatalf("runWithTimeout() error = %v", err)
	}
}
