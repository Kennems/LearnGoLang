package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 切片容量的追加
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 切片的截取
	numbers = numbers[1:3]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	s := []int{1, 2, 3}
	s1 := s[0:2]
	fmt.Println("s1 = ", s1)
	s2 := make([]int, 3)
	copy(s2, s)

	s1[0] = 100
	fmt.Println("s = ", s)
	fmt.Println("s1 = ", s1)

	// s2是深拷贝所以修改s1不影响s2
	fmt.Println("s2 = ", s2)
}
