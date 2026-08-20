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
		Title:  "Can't Hurt Me",
		Author: "David Goggins",
		Copies: 5,
		ID:     "ABC01",
	},
	{
		Title:  "The Mountain is You",
		Author: "Briana Weist",
		Copies: 9,
		ID:     "ABC02",
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
