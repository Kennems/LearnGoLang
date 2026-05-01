package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	//fmt.Println(x)
	//b, ok := x.(int)
	//var num = 10 + b
	//fmt.Println(num, b, ok)

	// 反射实现这个功能
	//v := reflect.ValueOf(x)
	//fmt.Println(v)
	//var n = v + 12
	//fmt.Println(n)

	v := reflect.ValueOf(x)
	var n = v.Int() + 12
	fmt.Println(n)
}

func main() {
	var a = 13
	reflectValue(a)

}
