package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
	ID     string
}

var catalog = []Book{
	{
		Title:  "Never Finished",
		Author: "David Goggins",
		Copies: 5,
		ID:     "ABC03",
	},
	{
		Title:  "How to Win Friends and Influence People",
		Author: "Dale Carnegie",
		Copies: 9,
		ID:     "ABC04",
	},
}

func GetAllBooks() []Book {
	return catalog
}

func BookToString(book Book) string {
	result := fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
	return result
}

func GetBook(ID string) (Book, bool) {
	for _, book := range catalog {
		if book.ID == ID {
			return book, true
		}
	}

	return Book{}, false
}
