package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func reflectSetValue1(x interface{}) {
	//*x = 100

	//v, _ := x.(*int64)
	//*v = 120

	v := reflect.ValueOf(x)
	fmt.Println(v.Kind(), reflect.Int64)
	fmt.Println(v.Elem().Kind(), reflect.Int64)
	if v.Elem().Kind() == reflect.Ptr {
		v.SetPointer(unsafe.Pointer(&v))
	} else if v.Kind() == reflect.Int64 {
		v.Elem().SetInt(120)
	}
}

func main() {
	var a int64 = 100
	reflectSetValue1(&a)
	fmt.Println(a)
}
