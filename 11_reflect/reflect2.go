package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user call")
	fmt.Printf("%v\n", this)
}

//func main() {
//	user := User{1, "Kennem", 20}
//	DoFileAndMethod(user)
//}

func DoFileAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	// 获取input的value
	inputValue := reflect.ValueOf(input)

	fmt.Println("inputType = ", inputType.Name())
	fmt.Println("inputValue = ", inputValue)

	// 通过type获取里面的字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i += 1 {
		m := inputType.Method(i)
		fmt.Printf("%s : %v\n", m.Name, m.Type)
	}

}
