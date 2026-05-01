package main

import "unsafe"

func makeSlice() []int {
	return make([]int, 3)
}

func newInt() *int {
	return new(int)
}

func emptyStructSize() uintptr {
	return unsafe.Sizeof(struct{}{})
}

func buildSet(items []string) map[string]struct{} {
	set := make(map[string]struct{}, len(items))
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

func typedNilInInterface() any {
	var p *int
	return p
}
