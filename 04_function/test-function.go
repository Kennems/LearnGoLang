package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	d := 200

	return c, d
}

// 返回多个返回值，有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---foo3---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 给有名称的返回值变量赋值， 初始化的默认值是0
	// r1, r2 作用域空间 是foo3 整个函数体的{}空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	r1 = 1000
	r2 = 2000
	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---foo4---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("abc", 222)
	fmt.Println("c = ", c)

	res1, res2 := foo2("abc", 222)
	fmt.Println("res1 = ", res1, "res2", res2)

	res3, res4 := foo3("abc", 222)
	fmt.Println("res3 = ", res3, "res4", res4)

	res5, res6 := foo4("abc", 222)
	fmt.Println("res5 = ", res5, "res6", res6)
}
