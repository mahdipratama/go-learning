package books_test

import (
	"books"
	"testing"
)

func TestBookToString_FormatBookInfoAsString(t *testing.T) {
	input := books.Book{
		Title:  "Can't Hurt Me",
		Author: "David Goggins",
		Copies: 5,
	}

	want := "Can't Hurt Me by David Goggins (copies: 5)"
	got := books.BookToString(input)
	if want != got {
		// instead of %v (print any value), we've use %q to specific values in strings
		t.Fatalf("want: %q \n got: %q", want, got)
	}
}
