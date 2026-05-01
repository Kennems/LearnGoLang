package main

import "fmt"

func appendOne(numbers []int) []int {
	return append(numbers, 1)
}

func trimNumbers(numbers []int) []int {
	return numbers[1:3]
}

func cloneSlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = appendOne(numbers)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = trimNumbers(numbers)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	s := []int{1, 2, 3}
	s1 := s[0:2]
	fmt.Println("s1 = ", s1)
	s2 := cloneSlice(s)

	s1[0] = 100
	fmt.Println("s = ", s)
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)
}
