package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumbers(5)
	go printNumbers(3)

	time.Sleep(3 * time.Second)
	fmt.Println("Main function finished")
}

func printNumbers(n int) {
	for i := 0; i < n; i += 1 {
		fmt.Println(i)
		time.Sleep(200 * time.Millisecond)
	}
}
