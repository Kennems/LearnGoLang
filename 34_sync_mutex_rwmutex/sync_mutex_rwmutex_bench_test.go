package main

import "testing"

func BenchmarkMutexCounterInc(b *testing.B) {
	var counter MutexCounter
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Inc()
		}
	})
}

func BenchmarkAtomicIncrement(b *testing.B) {
	var n int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomicIncrement(&n)
		}
	})
}
