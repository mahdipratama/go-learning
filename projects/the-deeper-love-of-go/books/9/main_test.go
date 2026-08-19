package main

import (
	"testing"
)

func TestBookToString_FormatBookInfoAsString(t *testing.T) {
	input := Book{
		title:  "Can't Hurt Me",
		author: "David Goggins",
		copies: 5,
	}

	want := "Can't Hurt Me by David Goggins (copies: 5)"
	got := BookToString(input)
	if want != got {
		// instead of %v (print any value), we've use %q to specific values in strings
		t.Fatalf("want: %q \n got: %q", want, got)
	}
}
