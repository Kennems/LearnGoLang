package main

import "fmt"

type Number interface {
	int | int32 | float64
}

// 定义一个getData的泛型方法
func getData[T any](value T) T {
	return value
}

func Add[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1, 2.1))
}
