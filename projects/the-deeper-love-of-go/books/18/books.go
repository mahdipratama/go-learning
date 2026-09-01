package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
	"sync"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     string
}

type Catalog struct {
	// Every method that accesses the catalog now needs to be mutex-aware
	// need to get the lock before trying to read the map data
	mu   *sync.RWMutex
	data map[string]Book
}

func (catalog Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(catalog))
}

func (book Book) String() string {
	result := fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
	return result
}

func (book *Book) SetCopies(copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	book.Copies = copies
	return nil
}

func (catalog Catalog) GetCopies(ID string) (int, error) {
	book, ok := catalog[ID]
	if !ok {
		return 0, fmt.Errorf("ID %q not found", ID)
	}

	return book.Copies, nil
}

func (catalog Catalog) GetBook(ID string) (Book, bool) {
	book, ok := catalog[ID]

	return book, ok
}

func (catalog Catalog) AddBook(book Book) error {
	_, ok := catalog[book.ID]
	if ok {
		return fmt.Errorf("ID: %q already exists", book.ID)
	}

	catalog[book.ID] = book

	return nil
}

func (catalog Catalog) SetCopies(ID string, copies int) error {

	book, ok := catalog.GetBook(ID)
	if !ok {
		return fmt.Errorf("ID: %q not found", ID)
	}

	err := book.SetCopies(copies)
	if err != nil {
		return err
	}

	catalog[ID] = book

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
