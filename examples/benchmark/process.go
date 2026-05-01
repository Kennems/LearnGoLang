package benchmark

func Process(data []int) {
	for i := 0; i < len(data); i++ {
		data[i] = data[i] * data[i]
	}
}
