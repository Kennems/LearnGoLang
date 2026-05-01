package main

import "fmt"

func main() {
	a, b := 10, 20
	swap(&a, &b)
	//fmt.Println("a = ", a, ", b = ", b)

	var p *int
	p = &a
	fmt.Println("p = ", p)
	fmt.Println("&a = ", &a)

	var pp **int
	pp = &p
	fmt.Println("pp = ", pp)
}

func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
