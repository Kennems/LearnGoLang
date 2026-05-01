package main

import (
	"testing"
	"time"
)

func TestOnceCounter(t *testing.T) {
	var c onceCounter

	for i := 0; i < 10; i++ {
		if got := c.Init(); got != 1 {
			t.Fatalf("c.Init() = %d, want 1", got)
		}
	}
}

func TestSquareWithWaitGroup(t *testing.T) {
	got := squareWithWaitGroup([]int{1, 2, 3, 4})
	want := []int{1, 4, 9, 16}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("squareWithWaitGroup() = %v, want %v", got, want)
		}
	}
}

func TestSignalBox(t *testing.T) {
	box := newSignalBox()
	done := make(chan string, 1)

	go func() {
		done <- box.Wait()
	}()

	time.Sleep(20 * time.Millisecond)
	box.Set("ready")

	select {
	case got := <-done:
		if got != "ready" {
			t.Fatalf("Wait() = %q, want %q", got, "ready")
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for cond signal")
	}
}
