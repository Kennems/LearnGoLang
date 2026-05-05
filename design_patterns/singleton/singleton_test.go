package singleton

import (
	"sync"
	"testing"
)

const goroutineCount = 1000

// 单线程测试
func TestSingletonSingleThread(t *testing.T) {
	a := GetSingleton()
	b := GetSingleton()
	if a != b {
		t.Fatalf("Singleton failed: a != b")
	}
}

// 并发测试
func TestSingletonConcurrent(t *testing.T) {
	start := make(chan struct{})
	instances := [goroutineCount]Singleton{}
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go func(idx int) {
			<-start
			instances[idx] = GetSingleton()
			wg.Done()
		}(i)
	}

	// 同时启动所有 goroutine
	close(start)
	wg.Wait()

	// 验证所有实例都相等
	for i := 0; i < goroutineCount-1; i++ {
		if instances[i] != instances[i+1] {
			t.Fatalf("Singleton failed: instances[%d] != instances[%d]", i, i+1)
		}
	}
}
