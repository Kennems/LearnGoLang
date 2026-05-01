package main

import "testing"

func TestCaptureMemSnapshot(t *testing.T) {
	s := captureMemSnapshot()
	if s.NumGoroutine == 0 {
		t.Fatal("NumGoroutine should be positive")
	}
}

func TestForceGC(t *testing.T) {
	if n := forceGC(); n == 0 {
		t.Fatal("NumGC should be positive after forcing GC in normal runtime")
	}
}
