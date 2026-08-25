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

type Catalog map[string]Book

func (catalog Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(catalog))
}

func (book Book) String() string {
	result := fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
	return result
}

func (catalog Catalog) GetBook(ID string) (Book, bool) {
	book, ok := catalog[ID]

	return book, ok
}

func (catalog Catalog) AddBook(book Book) {
	catalog[book.ID] = book
}

func (book *Book) SetCopies(copies int) {
	// Below, is a shallow copy: not exactly updating the book.Copies
	// fmt.Println("Before update book.Copies = ", book.Copies)
	// book.Copies = copies
	// fmt.Println("After update book.Copies = ", book.Copies)
	book.Copies = copies

}

func GetCatalog() Catalog {
	return Catalog{
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
