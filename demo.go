package main

import (
	"fmt"
	"time"
)

func goFunc(i int) {
	fmt.Println("goroutine", i, "...")
}

func main() { // 函数左括号一定和函数名在同一行
	for i := 0; i < 1000; i += 1 {
		go goFunc(i)
	}

	time.Sleep(time.Second)
}
