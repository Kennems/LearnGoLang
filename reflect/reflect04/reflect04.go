package main

import (
	"fmt"
	"reflect"
)

func reflectFn(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Println("type is invalid")
	case reflect.Int:
		fmt.Printf("type is int %v\n", v.Int())
	case reflect.Float32:
		fmt.Printf("type is float32 %v\n", v.Float())
	case reflect.Float64:
		fmt.Printf("type is float64 %v\n", v.Float())
	case reflect.String:
		fmt.Printf("type is string %v\n", v.String())
	default:
		fmt.Printf("type is %v\n", v.Type())
	}
}

func main() {
	var a float32 = 3.12
	var b = 100
	var c = "hello"

	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
}
