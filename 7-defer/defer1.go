package main

import "fmt"

// defer的应用场景：1、文件操作 2、互斥锁 3、数据库连接
func func1() {
	fmt.Println("A")
}
func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

/*
*
defer的执行顺序
*/
//func main() {
//	defer fmt.Println("main end1")
//	defer fmt.Println("main end2")
//
//	fmt.Println("main : hello go 1")
//	fmt.Println("main : hello go 2")
//
//	defer func1()
//	defer func2()
//	defer func3()
//	/*
//		main : hello go 1
//		main : hello go 2
//		C
//		B
//		A
//		main end2
//		main end1
//	*/
//}
