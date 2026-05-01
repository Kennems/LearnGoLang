package main

import (
	"strings"
	"testing"
)

func TestConcatWithBuilder(t *testing.T) {
	got := concatWithBuilder([]string{"go", "-", "lang"})
	if got != "go-lang" {
		t.Fatalf("concatWithBuilder() = %q, want %q", got, "go-lang")
	}
}

func TestRuneCount(t *testing.T) {
	if got := runeCount("hello 你好"); got != 8 {
		t.Fatalf("runeCount() = %d, want 8", got)
	}
}

func TestFormatExamples(t *testing.T) {
	v, plus, sharp := formatExamples(formatStudent{ID: 1, Name: "Ken"})
	if v == plus || plus == sharp {
		t.Fatal("expected different format outputs")
	}
	if !strings.Contains(plus, "ID:1") {
		t.Fatalf("unexpected %%+v output: %q", plus)
	}
}
