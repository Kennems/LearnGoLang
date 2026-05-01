package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	name string
	age  int
}

// 反射是指程序在运行期间对程序本身访问和修改的能力
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
}

func main() {
	a := 10
	b := 23.99
	c := true
	d := "hello"
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e myInt = 10
	var p = Person{
		"ken",
		22,
	}
	reflectFn(e)

	reflectFn(p)

	var h = 25
	reflectFn(&h)

}
