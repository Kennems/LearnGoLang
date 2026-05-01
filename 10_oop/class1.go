package main

import "fmt"

// 类首字母大写表示包外可以访问
type Hero struct {
	// 属性首字母大写表示类外可以访问
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	fmt.Println("hero = ", this.Name, this.Ad, this.Level)
}

/*
this 只是当前Hero的一个拷贝
func (this Hero) GetName() { fmt.Println("Name = ", this.Name)
}
func (this Hero) SetName(newName string) {
	this.Name = newName
}
*/
// 方法同理，大学表示包外可以访问
func (this *Hero) GetName() {
	fmt.Println("Name = ", this.Name)
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

//func main() {
//	hero := Hero{Name: "xx", Ad: 100, Level: 1}
//	hero.GetName()
//	hero.Show()
//	hero.SetName("Kennem")
//	hero.Show()
//}
