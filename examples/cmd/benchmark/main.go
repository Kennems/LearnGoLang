package main

import (
	"fmt"

	"go_study_examples/benchmark"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	benchmark.Process(data)
	fmt.Println(data)
}
