package main

import (
	"fmt"
	"net/http"
)

var message = "Hello"

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(hello))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
