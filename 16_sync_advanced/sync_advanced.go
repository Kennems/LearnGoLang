package main

import (
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu sync.RWMutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *Counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.n
}

func initOnce(once *sync.Once, fn func()) {
	once.Do(fn)
}

func atomicInc(n *int64) int64 {
	return atomic.AddInt64(n, 1)
}
