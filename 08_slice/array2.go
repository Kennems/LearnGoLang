package main

import "fmt"

func printArray2(myArray []int) {
	for _, value := range myArray {
		fmt.Println("value =", value)
	}
	myArray[0] = 100
}

//func main() {
//	// 引用传递
//	myArray := []int{1, 2, 3, 4} //动态数组，切片slice
//
//	fmt.Printf("myArray type is %T", myArray)
//	fmt.Println()
//
//	printArray2(myArray)
//
//	fmt.Println("===")
//
//	for _, value := range myArray {
//		fmt.Println("value: ", value)
//	}
//}
