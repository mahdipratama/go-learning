package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

var catalog []Book

func GetAllBooks() []Book {

	catalog = []Book{
		{
			Title:  "Can't Hurt Me",
			Author: "David Goggins",
			Copies: 5,
		},
		{
			Title:  "The Mountain is You",
			Author: "Briana Weist",
			Copies: 9,
		},
	}

	return catalog

}

func BookToString(book Book) string {
	result := fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
	return result
}
