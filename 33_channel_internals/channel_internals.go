package main

import "time"

func syncOverUnbuffered(v int) int {
	ch := make(chan int)
	go func() {
		ch <- v
	}()
	return <-ch
}

func bufferedChannelState() (int, int) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	return len(ch), cap(ch)
}

func readClosedBufferedChannel() (int, bool, int, bool) {
	ch := make(chan int, 1)
	ch <- 18
	close(ch)
	v1, ok1 := <-ch
	v2, ok2 := <-ch
	return v1, ok1, v2, ok2
}

func sendOnClosedChannelPanics() (panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()

	ch := make(chan int)
	close(ch)
	ch <- 1
	return false
}

func receiveWithSelectTimeout(ch <-chan int, timeout time.Duration) bool {
	select {
	case <-ch:
		return true
	case <-time.After(timeout):
		return false
	}
}
