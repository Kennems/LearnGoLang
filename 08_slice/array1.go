package main

import "fmt"

func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index =", index, "value =", value)
	}
}

//func main() {
//	// 固定长度的数组
//	var array1 = [5]int{1, 2, 3, 4, 5}
//
//	array2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	array3 := [4]int{11, 22, 33, 44}
//
//	for i := 0; i < len(array1); i++ {
//		fmt.Printf("%d ", array1[i])
//	}
//	fmt.Println()
//
//	for i := 0; i < len(array1); i += 1 {
//		fmt.Printf("%d ", array1[i])
//	}
//	fmt.Println()
//
//	for index, value := range array1 {
//		fmt.Println("index :", index, "value :", value)
//	}
//
//	fmt.Printf("myArray1: types = %T, value = %v\n", array1, array1)
//	fmt.Printf("myArray2: types = %T, value = %v\n", array2, array2)
//	fmt.Printf("myArray3: types = %T, value = %v\n", array3, array3)
//
//	printArray(array3)
//	// 静态数组只能传入对应长度的数组
//	// printArray(array2)
//	// 并且函数中修改不了数组中的值
//}
