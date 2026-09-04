package main

import (
	"fmt"
	"net/http"
)

var message = "Hello"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/{name}", hello)
	mux.HandleFunc("/goodbye", goodbye)

	http.ListenAndServe("localhost:3000", mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintf(w, "Hello %s\n", name)
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Goodbye world!")
}
