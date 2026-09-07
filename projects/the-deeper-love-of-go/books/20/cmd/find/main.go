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

	ID := os.Args[1]
	addr := books.NewClient("localhost:3000")

	book, err := addr.GetBook(ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(book)
}
