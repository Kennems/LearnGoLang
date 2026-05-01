package main

import "sort"

func buildIntMap() map[int]string {
	return map[int]string{
		3: "c",
		1: "a",
		2: "b",
	}
}

func sortedMapKeys(m map[int]string) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func deleteDuringRange(m map[int]int) {
	for k := range m {
		if k%2 == 0 {
			delete(m, k)
		}
	}
}

func mapLookup[K comparable, V any](m map[K]V, key K) (V, bool) {
	v, ok := m[key]
	return v, ok
}

func mapLengthAfterDelete() int {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	delete(m, 2)
	return len(m)
}
