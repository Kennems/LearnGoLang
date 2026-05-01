package main

import "runtime"

type memSnapshot struct {
	NumGoroutine int
	HeapAlloc    uint64
	NumGC        uint32
}

func captureMemSnapshot() memSnapshot {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	return memSnapshot{
		NumGoroutine: runtime.NumGoroutine(),
		HeapAlloc:    ms.HeapAlloc,
		NumGC:        ms.NumGC,
	}
}

func forceGC() uint32 {
	before := captureMemSnapshot().NumGC
	runtime.GC()
	after := captureMemSnapshot().NumGC
	if after > before {
		return after
	}
	return before
}
