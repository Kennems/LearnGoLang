package main

import (
	"fmt"
	"net/http"

	pprofexample "go_study_examples/pprof"
)

func main() {
	http.HandleFunc("/", pprofexample.SayHelloName)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("ListenAndServe: %s", err)
	}
}
