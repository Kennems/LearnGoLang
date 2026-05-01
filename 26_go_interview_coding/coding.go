package main

import "sync"

func alternatingSequence(n int) []int {
	oddCh := make(chan int)
	evenCh := make(chan int)
	done := make(chan struct{})
	out := make([]int, 0, n)
	var mu sync.Mutex

	go func() {
		for i := 1; i <= n; i += 2 {
			<-oddCh
			mu.Lock()
			out = append(out, i)
			mu.Unlock()
			evenCh <- i + 1
		}
		close(done)
	}()

	go func() {
		for i := 2; i <= n; i += 2 {
			<-evenCh
			if i <= n {
				mu.Lock()
				out = append(out, i)
				mu.Unlock()
			}
			if i < n {
				oddCh <- i + 1
			}
		}
	}()

	oddCh <- 1
	<-done
	return out
}

func workerPoolCount(tasks int, workers int) int {
	ch := make(chan int, tasks)
	done := make(chan struct{}, tasks)

	for i := 0; i < workers; i++ {
		go func() {
			for range ch {
				done <- struct{}{}
			}
		}()
	}

	for i := 0; i < tasks; i++ {
		ch <- i
	}
	close(ch)

	count := 0
	for i := 0; i < tasks; i++ {
		<-done
		count++
	}
	return count
}
