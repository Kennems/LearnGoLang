package main

import (
	"reflect"
	"testing"
	"time"
)

func TestReceiveWithTimeout(t *testing.T) {
	ch := make(chan int)

	if _, ok := receiveWithTimeout(ch, 5*time.Millisecond); ok {
		t.Fatal("expected timeout")
	}
}

func TestFanIn(t *testing.T) {
	a := make(chan int, 1)
	b := make(chan int, 1)
	a <- 1
	b <- 2

	var got []int
	for v := range fanIn(a, b) {
		got = append(got, v)
	}

	if !reflect.DeepEqual(got, []int{1, 2}) && !reflect.DeepEqual(got, []int{2, 1}) {
		t.Fatalf("fanIn() = %v", got)
	}
}
