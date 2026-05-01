package main

import (
	"fmt"
	"time"
)

// 原始实现
//func findMaxDifference(nums []int) int {
//	maxDiff := 0
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			diff := nums[j] - nums[i]
//			if diff > maxDiff {
//				maxDiff = diff
//			}
//		}
//	}
//	return maxDiff
//}

func findMaxDifference(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	minVal := nums[0]
	maxDiff := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}

		diff := nums[i] - minVal
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return maxDiff
}

func main() {
	start := time.Now() // 获取当前时间
	n := make([]int, 100005)
	for i := 1; i < 100000; i++ {
		n[i] = i
	}
	fmt.Println(findMaxDifference(n))
	cost := time.Since(start) // 计算此时与start的时间差
	fmt.Println(cost)
}
