package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	// os.Args is a slice of strings, came from user input as an arguments then store it
	// os.Args[0]: Contains the name or path of the running program executable.
	// os.Args[1:]: Contains all the actual user-provided arguments passed to the program.

	// HOW TO USE:
	// Current Folder -> go run main.go [bookId]
	// root -> go run ./cmd/find [bookId]

	if len(os.Args) != 2 {
		fmt.Println("Usage: find <BOOK ID>")
		return
	}

	catalog, err := books.OpenCatalog("testdata/catalog")
	if err != nil {
		fmt.Printf("Opening catalog: %v \n", err)
		return
	}

	ID := os.Args[1]

	book, ok := catalog.GetBook(ID)
	if !ok {
		fmt.Println("Sorry, Couldn't find that book in the catalog")
		return
	}

	fmt.Println(book)

}
