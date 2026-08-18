package main

import "fmt"

// Declaring struct type
type Book struct {
	title  string
	author string
	copies int
}

func main() {
	// First phase of struct assignment
	// var book = Book{
	// 	title:  "Can't Hurt Me",
	// 	author: "David Goggins",
	// 	copies: 5,
	// }

	// Second phase of struct assigntment
	book := Book{ // "This “colon-equals” syntax is technically known as the short declaration form,
		title:  "Can't Hurt Me",
		author: "David Goggins",
		copies: 5,
	}

	fmt.Println("Books in stock")

	printBook(book)
}

func printBook(book Book) {
	fmt.Println(book.title, "by", book.author, "-", book.copies, "copies")

}
