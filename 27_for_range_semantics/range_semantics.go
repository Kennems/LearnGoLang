package main

type item struct {
	Value int
}

func collectElementAddresses(items []item) []*item {
	out := make([]*item, 0, len(items))
	for i := range items {
		out = append(out, &items[i])
	}
	return out
}

func collectRangeVariableAddresses(items []item) []*item {
	out := make([]*item, 0, len(items))
	for _, v := range items {
		out = append(out, &v)
	}
	return out
}

func incrementByRangeValue(items []item) {
	for _, v := range items {
		v.Value++
	}
}

func incrementByIndex(items []item) {
	for i := range items {
		items[i].Value++
	}
}

func captureLoopValues(values []int) []int {
	out := make(chan int, len(values))
	for _, v := range values {
		go func(n int) {
			out <- n
		}(v)
	}

	got := make([]int, 0, len(values))
	for range values {
		got = append(got, <-out)
	}
	return got
}
