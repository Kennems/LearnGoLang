package main

import "testing"

func TestHeroSetName(t *testing.T) {
	hero := Hero{Name: "old", Ad: 10, Level: 1}
	hero.SetName("new")

	if hero.Name != "new" {
		t.Fatalf("hero.Name = %q, want %q", hero.Name, "new")
	}
}

func TestChangeBookVsPointer(t *testing.T) {
	book := Book{title: "Go", auth: "kk"}

	changeBook(book)
	if book.auth != "kk" {
		t.Fatalf("changeBook modified original auth to %q", book.auth)
	}

	changeBook2(&book)
	if book.auth != "777" {
		t.Fatalf("changeBook2() auth = %q, want %q", book.auth, "777")
	}
}

func TestAnimalImplementations(t *testing.T) {
	cat := &Cat{color: "white"}
	dog := &Dog{color: "yellow"}

	if cat.GetType() != "Cat" || cat.GetColor() != "white" {
		t.Fatalf("unexpected cat values: type=%q color=%q", cat.GetType(), cat.GetColor())
	}
	if dog.GetType() != "Dog" || dog.GetColor() != "yellow" {
		t.Fatalf("unexpected dog values: type=%q color=%q", dog.GetType(), dog.GetColor())
	}
}

func TestSuperManEmbedsHuman(t *testing.T) {
	s := SuperMan{Human: Human{name: "Ken", sex: "Male"}, level: 100}

	if s.name != "Ken" || s.sex != "Male" || s.level != 100 {
		t.Fatalf("unexpected embedded values: %+v", s)
	}
}
