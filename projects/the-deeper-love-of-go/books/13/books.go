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

func GetAllBooks() []Book {
	return slices.Collect(maps.Values(catalog))
}

func BookToString(book Book) string {
	result := fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
	return result
}

func GetBook(ID string) (Book, bool) {
	book, ok := catalog[ID]

	return book, ok
}
