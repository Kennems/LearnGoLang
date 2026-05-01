package main

import (
	"fmt"
	"reflect"
)

type taggedUser struct {
	ID   int    `json:"id" db:"user_id"`
	Name string `json:"name" db:"user_name"`
}

func formatDiff(v any) (string, string, string) {
	return fmt.Sprintf("%v", v), fmt.Sprintf("%+v", v), fmt.Sprintf("%#v", v)
}

func readTag(fieldName string, key string) string {
	field, ok := reflect.TypeOf(taggedUser{}).FieldByName(fieldName)
	if !ok {
		return ""
	}
	return field.Tag.Get(key)
}

func ignoreSecondReturn() int {
	n, _ := divmod(7, 3)
	return n
}

func divmod(a, b int) (int, int) {
	return a / b, a % b
}
