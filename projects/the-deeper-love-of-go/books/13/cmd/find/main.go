package main

import (
	"books"
	"fmt"
)

func main() {

	book, ok := books.GetBook("ABC03")
	if !ok {
		fmt.Println("Sorry, Couldn't find that book in the catalog")
		return
	}

	fmt.Println(books.BookToString(book))

}
