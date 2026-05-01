package main

import "testing"

func TestMakeSlice(t *testing.T) {
	got := makeSlice()
	if len(got) != 3 {
		t.Fatalf("len(makeSlice()) = %d, want 3", len(got))
	}
}

func TestNewInt(t *testing.T) {
	if v := *newInt(); v != 0 {
		t.Fatalf("*newInt() = %d, want 0", v)
	}
}

func TestEmptyStructSize(t *testing.T) {
	if size := emptyStructSize(); size != 0 {
		t.Fatalf("emptyStructSize() = %d, want 0", size)
	}
}

func TestBuildSet(t *testing.T) {
	set := buildSet([]string{"A", "A", "B"})
	if len(set) != 2 {
		t.Fatalf("len(set) = %d, want 2", len(set))
	}
}

func TestTypedNilInInterface(t *testing.T) {
	if typedNilInInterface() == nil {
		t.Fatal("typed nil stored in interface should not equal nil")
	}
}
