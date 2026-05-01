package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMovieJSONTags(t *testing.T) {
	movie := Movie{
		Title:  "喜剧之王",
		Year:   2000,
		Price:  10,
		Actors: []string{"xingye", "haha"},
	}

	data, err := json.Marshal(movie)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if _, ok := got["rmb"]; !ok {
		t.Fatalf("json tags missing rmb field: %v", got)
	}
	if _, ok := got["price"]; ok {
		t.Fatalf("unexpected price field in json: %v", got)
	}
}

func TestBook3ImplementsReaderAndWriter(t *testing.T) {
	book := &Book3{name: "Go"}

	var r Reader = book
	if _, ok := r.(Writer); !ok {
		t.Fatal("Reader value should also satisfy Writer")
	}
}

func TestFindTagValues(t *testing.T) {
	tp := reflect.TypeOf(resume{})

	if got := tp.Field(0).Tag.Get("info"); got != "name" {
		t.Fatalf("first info tag = %q, want %q", got, "name")
	}
	if got := tp.Field(1).Tag.Get("info"); got != "sex" {
		t.Fatalf("second info tag = %q, want %q", got, "sex")
	}
}

func TestUserReflectionMetadata(t *testing.T) {
	tp := reflect.TypeOf(User{})

	if tp.Name() != "User" {
		t.Fatalf("type name = %q, want %q", tp.Name(), "User")
	}
	if tp.NumField() != 3 {
		t.Fatalf("NumField() = %d, want 3", tp.NumField())
	}
}
