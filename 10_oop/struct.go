package main

// 声明一种行的数据类型 myint, 是 int 的别名
type myInt int

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	book.auth = "666"
}

func changeBook2(book *Book) {
	// 指针传递
	book.auth = "777"
}

//func main() {
//	var a myInt = 1
//	fmt.Printf("a = %d, type of a = %T\n", a, a)
//	var b myInt = 2
//	fmt.Println(a + b)
//
//	var book1 Book
//	book1.title = "Golang"
//	book1.auth = "KK"
//
//	changeBook(book1)
//	fmt.Printf("%v\n", book1)
//
//	changeBook2(&book1)
//	fmt.Printf("%v\n", book1)
//}
