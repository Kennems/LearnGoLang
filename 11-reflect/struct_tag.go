package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"My Name"` // 后面为标签，解释属性的含义
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i += 1 {
		tagString := t.Field(i).Tag.Get("info")
		fmt.Println("info:", tagString)
	}
}

//func main() {
//	var re = resume{"Tom", "Male"}
//
//	findTag(&re)
//}
