package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i += 1 {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
