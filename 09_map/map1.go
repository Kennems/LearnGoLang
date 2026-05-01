package main

import "fmt"

func printMap(myMap map[string]string) {
	// cityMap是一个引用传递
	for key, value := range myMap {
		fmt.Println("key: ", key, ", value: ", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	// 声明myMap1 是一种map类型，key是string， value是string类型
	var myMap1 map[string]string

	if myMap1 != nil {
		fmt.Println("myMap1", myMap1)
	}

	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"

	fmt.Println("myMap1", myMap1)

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "Python"
	fmt.Println("myMap2", myMap2)

	myMap3 := map[string]string{
		"one":   "php",
		"two":   "java",
		"three": "C",
	}

	fmt.Println("myMap3", myMap3)

	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "Washington"

	for key, value := range cityMap {
		fmt.Println("key: ", key, "value: ", value)
	}

	delete(cityMap, "China")
	cityMap["USA"] = "DC"

	fmt.Println("-----")
	printMap(cityMap)
}
