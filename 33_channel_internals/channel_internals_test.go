package main

import (
	"testing"
	"time"
)

func TestSyncOverUnbuffered(t *testing.T) {
	if got := syncOverUnbuffered(7); got != 7 {
		t.Fatalf("syncOverUnbuffered() = %d, want 7", got)
	}
}

func TestBufferedChannelState(t *testing.T) {
	length, capacity := bufferedChannelState()
	if length != 2 || capacity != 3 {
		t.Fatalf("len/cap = %d/%d, want 2/3", length, capacity)
	}
}

func TestReadClosedBufferedChannel(t *testing.T) {
	v1, ok1, v2, ok2 := readClosedBufferedChannel()
	if v1 != 18 || !ok1 || v2 != 0 || ok2 {
		t.Fatalf("unexpected values: %d %v %d %v", v1, ok1, v2, ok2)
	}
}

func TestSendOnClosedChannelPanics(t *testing.T) {
	if !sendOnClosedChannelPanics() {
		t.Fatal("expected panic")
	}
}

func TestReceiveWithSelectTimeout(t *testing.T) {
	ch := make(chan int)
	if receiveWithSelectTimeout(ch, 5*time.Millisecond) {
		t.Fatal("expected timeout")
	}
}
