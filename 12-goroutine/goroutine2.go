package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			return
			//runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	go func(a int, b int) {
		fmt.Println("a = ", a, "b = ", b)
		return
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
