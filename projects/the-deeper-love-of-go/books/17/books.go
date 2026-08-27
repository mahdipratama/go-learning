package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
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

func (book *Book) SetCopies(copies int) error {
	// Below, is a shallow copy: not exactly updating the book.Copies
	// fmt.Println("Before update book.Copies = ", book.Copies)
	// book.Copies = copies
	// fmt.Println("After update book.Copies = ", book.Copies)

	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}

	book.Copies = copies

	return nil

}

func OpenCatalog(path string) (Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	catalog := Catalog{}
	err = json.NewDecoder(file).Decode(&catalog)

	if err != nil {
		return nil, err
	}

	return catalog, nil
}

func (catalog Catalog) Sync(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()
	err = json.NewEncoder(file).Encode(catalog)
	if err != nil {
		return err
	}

	return nil
}
