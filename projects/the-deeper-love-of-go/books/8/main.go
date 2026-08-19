package main

import "fmt"

// Declaring struct type
type Book struct {
	title  string
	author string
	copies int
}

func main() {
	TestBookToString_FormatBookInfoAsString()
	fmt.Println("It's all good!")
}

func BookToString(book Book) string {
	// Printf = takes a format string with placeholders
	// Spintf = same as Printf but doesn't print anything
	result := fmt.Sprintf("%v by %v - %v copies", book.title, book.author, book.copies)
	return result
}

func TestBookToString_FormatBookInfoAsString() {
	input := Book{ // "This “colon-equals” syntax is technically known as the short declaration form,
		title:  "Can't Hurt Me",
		author: "David Goggins",
		copies: 5,
	}

	want := "Can't Hurt Me by David Goggins - 5 copies"
	got := BookToString(input)
	if want != got {
		panic("BookToString: Wrong result!")
	}
}
