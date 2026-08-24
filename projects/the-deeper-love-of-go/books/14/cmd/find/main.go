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
	// root -> go run cmd/find [bookId]

	catalog := books.GetCatalog()

	if len(os.Args) != 2 {
		fmt.Println("Usage: find <BOOK ID>")
		return
	}

	ID := os.Args[1]

	book, ok := books.GetBook(catalog, ID)
	if !ok {
		fmt.Println("Sorry, Couldn't find that book in the catalog")
		return
	}

	fmt.Println(books.BookToString(book))

}

func GetCataloq() map[string]books.Book {
	return map[string]books.Book{
		"ABC03": {
			Title:  "Never Finished",
			Author: "David Goggins",
			Copies: 2,
			ID:     "ABC03",
		},
		"ABC04": {
			Title:  "The Mountain is You",
			Author: "Briana Weist",
			Copies: 1,
			ID:     "ABC04",
		},
	}
}
