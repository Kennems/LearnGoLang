package main

import "testing"

func TestFormatWorkerArgs(t *testing.T) {
	if got := formatWorkerArgs(10, 20); got != "a =  10 b =  20" {
		t.Fatalf("formatWorkerArgs() = %q", got)
	}
}
