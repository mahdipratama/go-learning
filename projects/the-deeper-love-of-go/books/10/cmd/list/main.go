package main

import (
	"books"
	"fmt"
)

func main() {
	book := books.Book{
		Title:  "The Mountain is You",
		Author: "Briana Weist",
		Copies: 4,
	}
	fmt.Println(books.BookToString(book))
}
