package main

import (
	"fmt"
	"net/http"
)

var message = "Hello"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/goodbye", goodbye)

	http.ListenAndServe("localhost:3000", mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Goodbye world!")
}
