package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // SuperMan类继承了Human类的方法
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

//func main() {
//	h := Human{"Kennem", "Male"}
//
//	h.Eat()
//	h.Walk()
//
//	//s := SuperMan{Human{"Kennem", "Male"}, 100}
//	var s SuperMan
//	s.name = "Kennem"
//	s.sex = "Male"
//	s.level = 100
//
//	s.Print()
//	s.Eat()
//	s.Fly()
//	s.Walk()
//}
