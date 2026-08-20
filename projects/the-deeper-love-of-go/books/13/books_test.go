package books_test

import (
	"books"
	"slices"
	"testing"
)

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	want := []books.Book{
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

	got := books.GetAllBooks()
	if !slices.Equal(want, got) {
		// Format placeholder "%#v": prints its corresponding data as a Go Value
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_FindsBookInCatalogByID(t *testing.T) {
	// Running all the test at the same time
	t.Parallel()

	want := books.Book{
		ID:     "ABC01",
		Title:  "Can't Hurt Me",
		Author: "David Goggins",
		Copies: 5,
	}

	got, ok := books.GetBook("ABC01")
	if !ok {
		t.Fatal("Book not found")
	}

	if want != got {
		t.Fatalf("want: %#v, got: %#v ", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()

	_, ok := books.GetBook("nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}
