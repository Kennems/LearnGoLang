package main

import (
	"sync"
	"sync/atomic"
)

type onceCounter struct {
	once sync.Once
	n    int32
}

func (c *onceCounter) Init() int32 {
	c.once.Do(func() {
		atomic.AddInt32(&c.n, 1)
	})
	return atomic.LoadInt32(&c.n)
}

func squareWithWaitGroup(nums []int) []int {
	results := make([]int, len(nums))
	var wg sync.WaitGroup

	for i, n := range nums {
		wg.Add(1)
		go func(idx, value int) {
			defer wg.Done()
			results[idx] = value * value
		}(i, n)
	}

	wg.Wait()
	return results
}

type signalBox struct {
	mu    sync.Mutex
	cond  *sync.Cond
	ready bool
	value string
}

func newSignalBox() *signalBox {
	box := &signalBox{}
	box.cond = sync.NewCond(&box.mu)
	return box
}

func (b *signalBox) Set(value string) {
	b.mu.Lock()
	b.value = value
	b.ready = true
	b.mu.Unlock()
	b.cond.Broadcast()
}

func (b *signalBox) Wait() string {
	b.mu.Lock()
	defer b.mu.Unlock()

	for !b.ready {
		b.cond.Wait()
	}
	return b.value
}
