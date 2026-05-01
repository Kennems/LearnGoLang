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
	fmt.Printf("类型：%v, 类型名称:%v 类型种类:%v \n", v, v.Name(), v.Kind())
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

	var i = [3]int{1, 2, 3}
	reflectFn(i)

	var j = []int{1, 2, 3}
	reflectFn(j)
	//类型：int, 类型名称:int 类型种类:int
	//类型：float64, 类型名称:float64 类型种类:float64
	//类型：bool, 类型名称:bool 类型种类:bool
	//类型：string, 类型名称:string 类型种类:string
	//类型：main.myInt, 类型名称:myInt 类型种类:int
	//类型：main.Person, 类型名称:Person 类型种类:struct
	//类型：*int, 类型名称: 类型种类:ptr
	//类型：[3]int, 类型名称: 类型种类:array
	//类型：[]int, 类型名称: 类型种类:slice
}
