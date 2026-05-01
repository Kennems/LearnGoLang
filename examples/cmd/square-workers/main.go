package main

import (
	"fmt"
	"time"

	"go_study_examples/concurrency"
)

func main() {
	start := time.Now()
	nums := []int{1, 2, 3, 4, 5}

	ch := make(chan int)
	for _, num := range nums {
		go func(n int) {
			ch <- concurrency.GetSquare(n)
		}(num)
	}

	for range nums {
		fmt.Println(<-ch)
	}

	fmt.Println(time.Since(start))
}
