package main

import (
	"testing"
	"time"
)

func TestReadAfterClose(t *testing.T) {
	v1, ok1, v2, ok2 := readAfterClose()
	if v1 != 18 || !ok1 || v2 != 0 || ok2 {
		t.Fatalf("unexpected closed channel reads: %d %v %d %v", v1, ok1, v2, ok2)
	}
}

func TestSendClosedChannelPanics(t *testing.T) {
	if !sendClosedChannelPanics() {
		t.Fatal("expected panic when sending on closed channel")
	}
}

func TestRecvWithTimeout(t *testing.T) {
	ch := make(chan int)
	if recvWithTimeout(ch, 5*time.Millisecond) {
		t.Fatal("expected timeout")
	}
}
