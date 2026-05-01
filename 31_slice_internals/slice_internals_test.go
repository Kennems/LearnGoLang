package main

import (
	"reflect"
	"testing"
)

func TestHeaderView(t *testing.T) {
	got := headerView(make([]int, 2, 5))
	if got.Len != 2 || got.Cap != 5 {
		t.Fatalf("headerView() = %+v", got)
	}
}

func TestSharedSubsliceMutatesBase(t *testing.T) {
	base := []int{0, 1, 2, 3}
	sub := sharedSubslice(base)
	sub[0] = 99

	if base[1] != 99 {
		t.Fatalf("base = %v, want shared backing array", base)
	}
}

func TestAppendWithoutGrowthSharesBackingArray(t *testing.T) {
	base := make([]int, 2, 4)
	base[0], base[1] = 1, 2

	sub := base[:2]
	appended := appendWithoutGrowth(sub, 3)
	appended[0] = 88

	if base[0] != 88 {
		t.Fatalf("base = %v, expected shared backing array", base)
	}
}

func TestAppendWithGrowthBreaksAlias(t *testing.T) {
	base := []int{1, 2}
	sub := base[:2:2]

	appended := appendWithGrowth(sub, 3, 4)
	appended[0] = 88

	if base[0] != 1 {
		t.Fatalf("base = %v, expected alias to be broken", base)
	}
}

func TestCloneSliceValues(t *testing.T) {
	src := []int{1, 2, 3}
	dst := cloneSliceValues(src)
	src[0] = 99

	if reflect.DeepEqual(src, dst) {
		t.Fatalf("dst unexpectedly aliases src: src=%v dst=%v", src, dst)
	}
	if !reflect.DeepEqual(dst, []int{1, 2, 3}) {
		t.Fatalf("dst = %v", dst)
	}
}
