package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     string
}

var catalog = map[string]Book{
	"ABC03": {
		Title:  "Never Finished",
		Author: "David Goggins",
		Copies: 5,
		ID:     "ABC03",
	},

	"ABC04": {
		Title:  "How to Win Friends and Influence People",
		Author: "Dale Carnegie",
		Copies: 9,
		ID:     "ABC04",
	},
}

func GetAllBooks() []Book {
	return slices.Collect(maps.Values(catalog))
}

func BookToString(book Book) string {
	result := fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
	return result
}

func GetBook(ID string) (Book, bool) {
	// with map, range gives us each key and element pair in turn.
	for _, book := range catalog {
		if book.ID == ID {
			return book, true
		}
	}

	return Book{}, false
}
