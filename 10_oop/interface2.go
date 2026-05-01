package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println(arg)

	// interface{} 区分引用的底层数据类型是什么
	// 给 interface() 提供 "类型断言" 机制

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)

		fmt.Printf("value type is %T\n", value)
	}
}

type Book1 struct {
	auth string
}

//func main() {
//	book := Book1{"Golang"}
//	myFunc(book)
//	myFunc(100)
//	myFunc("abcd")
//	myFunc(3.14)
//}
