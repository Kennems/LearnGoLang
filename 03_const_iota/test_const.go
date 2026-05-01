package main

import "fmt"

const (
	// 可以在const()中添加一个关键字iota, 每行的iota都会累加1， 第一行的iota的默认值是0
	//BEIJING  = 0
	//SHANGHAI = 1
	//NEWYORK  = 2

	BEIJING = 10 * iota
	SHANGHAI
	NEWYORK
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
	g, h = iota * 3, iota * 4
)

func main() {
	const length int = 100

	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("NEWYORK = ", NEWYORK)
	fmt.Println("NewYork = ", NEWYORK)
}
