package main

import (
	"sync"
	"sync/atomic"
)

type MutexCounter struct {
	mu sync.Mutex
	n  int64
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *MutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

type RWMutexCounter struct {
	mu sync.RWMutex
	n  int64
}

func (c *RWMutexCounter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *RWMutexCounter) Value() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.n
}

func atomicIncrement(n *int64) int64 {
	return atomic.AddInt64(n, 1)
}
