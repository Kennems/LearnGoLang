package main

import "testing"

func TestFindMaxDifference(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "empty", nums: nil, want: 0},
		{name: "single", nums: []int{9}, want: 0},
		{name: "descending", nums: []int{9, 7, 5, 3}, want: 0},
		{name: "increasing", nums: []int{1, 2, 6, 10}, want: 9},
		{name: "mixed", nums: []int{7, 1, 5, 3, 6, 4}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxDifference(tt.nums); got != tt.want {
				t.Fatalf("findMaxDifference(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
