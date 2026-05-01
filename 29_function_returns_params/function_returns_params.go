package main

func addOneByValue(n int) {
	n++
}

func mutateSliceElements(values []int) {
	for i := range values {
		values[i]++
	}
}

func appendValue(values []int, v int) []int {
	return append(values, v)
}

func appendValueByPointer(values *[]int, v int) {
	*values = append(*values, v)
}

func splitAndSum(a, b int) (int, int, int) {
	return a, b, a + b
}

func deferredNamedReturn() (result int) {
	result = 1
	defer func() {
		result++
	}()
	return result
}

func deferredUnnamedReturn() int {
	result := 1
	defer func() {
		result++
	}()
	return result
}
