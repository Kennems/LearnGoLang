package main

import "fmt"

var gA, gB int

// gC := 100 :=只能用在函数体内来声明
var gC = 100

func main() {
	var a int
	fmt.Println(a)

	var b int = 100
	fmt.Println("b = ", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	var c = 100
	fmt.Println("c = ", c)

	var cc = 1000
	fmt.Printf("cc = %d, type of cc = %T\n", cc, cc)

	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e : %T\n", e)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g : %T\n", g)

	fmt.Printf("gA = %d, gB = %d, gC = %d\n", gA, gB, gC)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Printf("xx = %d, yy = %d\n", xx, yy)

	var v1, v2 = 100, "abcd"
	fmt.Printf("v1 = ", v1, "v2 = ", v2)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Printf("vv = %d, jj = %t\n", vv, jj)
}
