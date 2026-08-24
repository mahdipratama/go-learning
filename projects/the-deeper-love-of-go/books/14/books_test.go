package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	catalog := getTestCataloq()

	want := []books.Book{
		{
			Title:  "The Mountain is You",
			Author: "Briana Weist",
			Copies: 1,
			ID:     "ABC04",
		},
		{
			Title:  "Never Finished",
			Author: "David Goggins",
			Copies: 2,
			ID:     "ABC03",
		},
	}

	got := books.GetAllBooks(catalog)
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})

	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_FindsBookInCatalogByID(t *testing.T) {
	// Running all the test at the same time
	t.Parallel()

	catalog := getTestCataloq()

	want := books.Book{
		Title:  "Never Finished",
		Author: "David Goggins",
		Copies: 2,
		ID:     "ABC03",
	}

	got, ok := books.GetBook(catalog, "ABC03")
	if !ok {
		t.Fatal("Book not found")
	}

	if want != got {
		t.Fatalf("want: %#v, got: %#v ", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()

	catalog := getTestCataloq()

	_, ok := books.GetBook(catalog, "nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddGivenBookToCatalog(t *testing.T) {
	t.Parallel()

	catalog := getTestCataloq()

	// Precondition
	_, ok := books.GetBook(catalog, "ABC03")
	if ok {
		t.Fatal("Book already present")
	}

	// Action
	books.AddBook(catalog, books.Book{
		ID:     "ABC03",
		Title:  "Never Finished",
		Author: "David Goggins",
		Copies: 2,
	})

	// Postcondition
	// _, ok := books.GetBook(catalog, "ABC03")
	if !ok {
		t.Fatalf("Added book not found")
	}

}

func getTestCataloq() map[string]books.Book {
	return map[string]books.Book{
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
