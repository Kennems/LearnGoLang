package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
)

type cache struct {
	data sync.Map
}

func (c *cache) LoadOrStore(key string, value string) string {
	actual, _ := c.data.LoadOrStore(key, value)
	return actual.(string)
}

func (c *cache) Load(key string) (string, bool) {
	v, ok := c.data.Load(key)
	if !ok {
		return "", false
	}
	return v.(string), true
}

type atomicState struct {
	value int32
}

func (s *atomicState) Inc() int32 {
	return atomic.AddInt32(&s.value, 1)
}

func (s *atomicState) SetIfZero() bool {
	return atomic.CompareAndSwapInt32(&s.value, 0, 1)
}

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func formatWithPool(prefix string, n int) string {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	fmt.Fprintf(buf, "%s:%d", prefix, n)
	return buf.String()
}
