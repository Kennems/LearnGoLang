package main

import "time"

func receiveWithTimeout(ch <-chan int, timeout time.Duration) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	case <-time.After(timeout):
		return 0, false
	}
}

func fanIn(a <-chan int, b <-chan int) <-chan int {
	out := make(chan int, 2)

	go func() {
		defer close(out)

		select {
		case v := <-a:
			out <- v
		case v := <-b:
			out <- v
		}

		select {
		case v := <-a:
			out <- v
		case v := <-b:
			out <- v
		}
	}()

	return out
}
