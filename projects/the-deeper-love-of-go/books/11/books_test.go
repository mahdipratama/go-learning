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
