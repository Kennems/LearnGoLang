package main

type sliceHeaderView struct {
	Len int
	Cap int
}

func headerView(s []int) sliceHeaderView {
	return sliceHeaderView{
		Len: len(s),
		Cap: cap(s),
	}
}

func sharedSubslice(base []int) []int {
	return base[1:3]
}

func appendWithoutGrowth(base []int, v int) []int {
	return append(base, v)
}

func appendWithGrowth(base []int, values ...int) []int {
	return append(base, values...)
}

func cloneSliceValues(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}
