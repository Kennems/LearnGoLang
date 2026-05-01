package main

import "unsafe"

type sample struct {
	A int32
	B int64
}

func fieldOffset() uintptr {
	return unsafe.Offsetof(sample{}.B)
}

func pointerRoundTrip(v *int) *int {
	p := unsafe.Pointer(v)
	u := uintptr(p)
	return (*int)(unsafe.Pointer(u))
}

func escapeCandidate() *int {
	v := 42
	return &v
}
