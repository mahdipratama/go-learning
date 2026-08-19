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
	return ""
}

func TestBookToString_FormatBookInfoAsString() {
	input := Book{ // "This “colon-equals” syntax is technically known as the short declaration form,
		title:  "Can't Hurt Me",
		author: "David Goggins",
		copies: 5,
	}

	want := "Can't Hurt Me - 5 copies"
	got := BookToString(input)
	if want != got {
		panic("BookToString: Wrong resul!")
	}
}
