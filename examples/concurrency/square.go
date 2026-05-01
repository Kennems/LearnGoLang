package concurrency

import "time"

func GetSquare(num int) int {
	time.Sleep(time.Second)
	return num * num
}
