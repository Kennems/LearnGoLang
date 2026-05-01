package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type formatStudent struct {
	ID   int
	Name string
}

func concatWithBuilder(parts []string) string {
	var b strings.Builder
	for _, part := range parts {
		b.WriteString(part)
	}
	return b.String()
}

func runeCount(s string) int {
	return utf8.RuneCountInString(s)
}

func formatExamples(v any) (string, string, string) {
	return fmt.Sprintf("%v", v), fmt.Sprintf("%+v", v), fmt.Sprintf("%#v", v)
}
