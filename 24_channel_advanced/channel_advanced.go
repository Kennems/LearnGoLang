package main

import "time"

func readAfterClose() (int, bool, int, bool) {
	ch := make(chan int, 1)
	ch <- 18
	close(ch)

	v1, ok1 := <-ch
	v2, ok2 := <-ch
	return v1, ok1, v2, ok2
}

func sendClosedChannelPanics() (panicked bool) {
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

func recvWithTimeout(ch <-chan int, d time.Duration) bool {
	select {
	case <-ch:
		return true
	case <-time.After(d):
		return false
	}
}
