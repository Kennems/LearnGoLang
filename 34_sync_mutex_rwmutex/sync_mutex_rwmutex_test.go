package main

import (
	"sync"
	"testing"
)

func TestMutexCounter(t *testing.T) {
	var counter MutexCounter
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()

	if got := counter.Value(); got != 100 {
		t.Fatalf("counter.Value() = %d, want 100", got)
	}
}

func TestRWMutexCounter(t *testing.T) {
	var counter RWMutexCounter
	counter.Inc()
	counter.Inc()

	if got := counter.Value(); got != 2 {
		t.Fatalf("counter.Value() = %d, want 2", got)
	}
}

func TestAtomicIncrement(t *testing.T) {
	var n int64
	if got := atomicIncrement(&n); got != 1 {
		t.Fatalf("atomicIncrement() = %d, want 1", got)
	}
}
