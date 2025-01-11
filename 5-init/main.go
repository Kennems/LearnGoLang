package main

import (
	// _ 表示导入lib1，但是不使用其中的接口
	_ "Demo/5-init/lib1"
	//. 表示不需要指定哪一个包，可以直接使用, 即可以直接使用方法名调用该包中的方法
	//. "Demo/5-init/lib2"

	// 以别名的方式
	mylib "Demo/5-init/lib2"
)

func main() {
	//lib1.Lib1Test()
	mylib.Lib2Test()
}
