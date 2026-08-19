package main

import "fmt"

type Book struct {
	title  string
	author string
	copies int
}

func main() {
	book := Book{
		title:  "The Mountain is You",
		author: "Briana Weist",
		copies: 4,
	}
	fmt.Println(BookToString(book))
}

func BookToString(book Book) string {
	result := fmt.Sprintf("%v by %v (copies: %v)", book.title, book.author, book.copies)
	return result
}
