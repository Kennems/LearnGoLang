package main

import (
	"strings"
	"testing"
)

func TestFormatDiff(t *testing.T) {
	v, plus, sharp := formatDiff(taggedUser{ID: 1, Name: "Ken"})
	if v == plus || plus == sharp {
		t.Fatal("expected different formatted outputs")
	}
	if !strings.Contains(plus, "ID:1") {
		t.Fatalf("unexpected %%+v output: %q", plus)
	}
	if !strings.Contains(sharp, "taggedUser") {
		t.Fatalf("unexpected %%#v output: %q", sharp)
	}
}

func TestReadTag(t *testing.T) {
	if got := readTag("ID", "json"); got != "id" {
		t.Fatalf("readTag(json) = %q, want %q", got, "id")
	}
	if got := readTag("Name", "db"); got != "user_name" {
		t.Fatalf("readTag(db) = %q, want %q", got, "user_name")
	}
}

func TestIgnoreSecondReturn(t *testing.T) {
	if got := ignoreSecondReturn(); got != 2 {
		t.Fatalf("ignoreSecondReturn() = %d, want 2", got)
	}
}
