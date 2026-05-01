package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var c Counter
	c.Inc()
	c.Inc()

	if c.Value() != 2 {
		t.Fatalf("c.Value() = %d, want 2", c.Value())
	}
}

func TestInitOnce(t *testing.T) {
	var once sync.Once
	count := 0

	initOnce(&once, func() { count++ })
	initOnce(&once, func() { count++ })

	if count != 1 {
		t.Fatalf("count = %d, want 1", count)
	}
}

func TestAtomicInc(t *testing.T) {
	var n int64
	if got := atomicInc(&n); got != 1 {
		t.Fatalf("atomicInc() = %d, want 1", got)
	}
}
