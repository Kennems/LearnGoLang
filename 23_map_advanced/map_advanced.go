package main

import "sort"

func sortedKeys(m map[int]string) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func deleteEvens(m map[int]int) {
	for k := range m {
		if k%2 == 0 {
			delete(m, k)
		}
	}
}
