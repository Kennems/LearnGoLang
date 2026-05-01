package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name" form:"username"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("Student name: %s, Age: %d, Score: %d", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) APrint() {
	fmt.Println("Print")
}

func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体")
	}

	field0 := t.Field(0)
	fmt.Println(field0)
	fmt.Println("字段名称：", field0.Name)
	fmt.Println("字段类型：", field0.Type)
	fmt.Println("字段Tag：", field0.Tag)
	fmt.Println("字段Tag：", field0.Tag.Get("json"))
	fmt.Println("字段Tag：", field0.Tag.Get("form"))
	//fmt.Println(field0.Name, field0.Tag, field0.Type)

	fmt.Println("-----------------------")

	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Println("字段名称：", field1.Name)
		fmt.Println("字段类型：", field1.Type)
		fmt.Println("字段Tag：", field1.Tag.Get("json"))
	}

	var fieldCount = t.NumField()
	fmt.Println("结构体有", fieldCount, "个属性")

	// 获取结构体属性对应的值
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))

	fmt.Println("-----------------------")

	for i := 0; i < fieldCount; i += 1 {
		fmt.Printf("属性名称:%v 属性值:%v 属性类型:%v 属性Tag:%v \n",
			t.Field(i).Name,
			t.Field(i),
			t.Field(i).Type,
			t.Field(i).Tag)
	}
}

func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct || t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体")
	}

	fmt.Println("-----------------------")

	method0 := t.Method(0)
	fmt.Println(method0.Name)
	fmt.Println(method0.Type)

	method1, ok := t.MethodByName("APrint")
	if ok {
		fmt.Println(method1.Name)
		fmt.Println(method1.Type)
	}
	fmt.Println("-----------------------")
	v.Method(0).Call(nil)
	v.MethodByName("APrint").Call(nil)
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)

	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(0))
	setInfoFn := v.MethodByName("SetInfo")
	fmt.Println(setInfoFn.Call(params))
	info = v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)

	fmt.Println("-----------------------")

	fmt.Println("方法数量", t.NumMethod())

}

func reflectChangeStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Ptr {
		fmt.Println("传入的不是结构体指针类型")
		return
	} else if t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的不是结构体指针类型")
		return
	}

	name := v.Elem().FieldByName("Name")
	name.SetString("Li")

	age := v.Elem().FieldByName("Age")
	age.SetInt(11)

}

func main() {
	s := Student{
		"Ken",
		22,
		100,
	}

	//PrintStructField(s)
	//PrintStructFn(&s)

	reflectChangeStruct(&s)
	fmt.Printf("%#v\n", s)
}
