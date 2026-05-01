package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book3 struct {
	name string
}

func (this *Book3) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book3) WriteBook() {
	fmt.Println("Write a Book")
}

//func main() {
//	// b: pair<type:Book3, value:book{}地址>
//	b := &Book3{}
//
//	// r: pair<type:, value:>
//	var r Reader
//	// r: pair<type:Book3, value:book{}地址>
//	r = b
//
//	r.ReadBook()
//
//	var w Writer
//	// r: pair<type:Book, value:book{}地址>
//	// 类型断言的语法是 x.(T)，如果 x 的实际类型是 T，那么这个断言就成功，并返回类型 T 的值。
//	// 如果 x 的实际类型不是 T，那么这个断言就失败，并返回类型 T 的零值。
//	// 因为 r 的实际类型是 Book3, 而Book3已经实现了Writer ，所以断言成功，返回 book{}地址
//	w = r.(Writer)
//
//	w.WriteBook()
//}
