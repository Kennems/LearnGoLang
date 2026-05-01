package main

import (
	"sync"
	"testing"
)

func TestCacheLoadOrStore(t *testing.T) {
	var c cache

	if got := c.LoadOrStore("go", "lang"); got != "lang" {
		t.Fatalf("LoadOrStore() = %q, want %q", got, "lang")
	}
	if got, ok := c.Load("go"); !ok || got != "lang" {
		t.Fatalf("Load() = %q, %v", got, ok)
	}
}

func TestAtomicState(t *testing.T) {
	var s atomicState
	if !s.SetIfZero() {
		t.Fatal("expected CAS from zero to succeed")
	}
	if s.SetIfZero() {
		t.Fatal("expected CAS from non-zero to fail")
	}
	if got := s.Inc(); got != 2 {
		t.Fatalf("Inc() = %d, want 2", got)
	}
}

func TestFormatWithPool(t *testing.T) {
	if got := formatWithPool("go", 36); got != "go:36" {
		t.Fatalf("formatWithPool() = %q, want %q", got, "go:36")
	}
}

func TestCacheConcurrentLoadOrStore(t *testing.T) {
	var c cache
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.LoadOrStore("key", "value")
		}()
	}
	wg.Wait()

	if got, ok := c.Load("key"); !ok || got != "value" {
		t.Fatalf("Load() = %q, %v", got, ok)
	}
}
