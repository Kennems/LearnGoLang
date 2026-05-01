package main

import (
	"time"
)

//func main() {
//	start := time.Now() // 获取当前时间
//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//
//	ch := make(chan int)
//	for _, num := range nums {
//		go func(n int) {
//			res := getSquare(n) // 调用getSquare函数计算n的平方
//			ch <- res           // 将计算结果发送到通道ch中
//		}(num)
//	}
//
//	for i := 0; i < len(nums); i++ {
//		fmt.Println(<-ch) // 从通道ch中接收计算结果并打印
//	}
//
//	cost := time.Since(start) // 计算此时与start的时间差
//	fmt.Println(cost)         // 打印程序运行所耗费的时间
//}

func getSquare(num int) int {
	time.Sleep(1 * time.Second) // 模拟一个耗时操作
	return num * num
}

//func main() {
//	start := time.Now() // 获取当前时间
//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//
//	for _, num := range nums {
//		fmt.Println(getSquare(num))
//	}
//	cost := time.Since(start) // 计算此时与start的时间差
//	fmt.Println(cost)
//}
//
//func getSquare(num int) int {
//	time.Sleep(1 * time.Second) // 模拟一个耗时操作
//	return num * num
//}
