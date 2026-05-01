package pprofexample

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func SayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "Hello World!")
}
