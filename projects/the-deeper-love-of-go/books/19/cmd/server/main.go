package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	// usage go run ./cmd/server testdata/catalog

	if len(os.Args) != 2 {
		fmt.Println("Usage: server <CATALOG FILE>")
		return
	}

	path := os.Args[1]

	catalog, err := books.OpenCatalog(path)
	if err != nil {
		fmt.Printf("Opening catalog: %v \n", err)
		return
	}

	err = books.ListenAndServe(":3000", catalog)
	if err != nil {
		fmt.Println(err)
	}
}
