package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	catalog := getTestCatalog()

	bookList := catalog.GetAllBooks()
	assertTestBooks(t, bookList)

}

func TestOpenCatalog_ReadSameDataWrittenBySync(t *testing.T) {
	t.Parallel()

	catalog := getTestCatalog()

	path := t.TempDir() + "/catalog"
	err := catalog.Sync(path)
	if err != nil {
		t.Fatal(err)
	}

	newCatalog, err := books.OpenCatalog("testdata/catalog.new")
	if err != nil {
		t.Fatal(err)
	}

	bookList := newCatalog.GetAllBooks()
	assertTestBooks(t, bookList)

}

func TestGetBook_FindsBookInCatalogByID(t *testing.T) {
	// Running all the test at the same time
	t.Parallel()

	catalog := getTestCatalog()

	want := books.Book{
		Title:  "Never Finished",
		Author: "David Goggins",
		Copies: 2,
		ID:     "ABC03",
	}

	got, ok := catalog.GetBook("ABC03")
	if !ok {
		t.Fatal("Book not found")
	}

	if want != got {
		t.Fatalf("want: %#v, got: %#v ", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()

	catalog := getTestCatalog()

	_, ok := catalog.GetBook("nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddGivenBookToCatalog(t *testing.T) {
	t.Parallel()

	catalog := getTestCatalog()

	// Precondition
	_, ok := catalog.GetBook("ABC05")
	if ok {
		t.Fatal("Book already present")
	}

	// Action
	err := catalog.AddBook(books.Book{
		ID:     "ABC05",
		Title:  "Atomic Habit",
		Author: "James Clear",
		Copies: 2,
	})

	if err != nil {
		t.Fatal(err)
	}

	// Postcondition
	_, ok = catalog.GetBook("ABC05")
	if !ok {
		t.Fatalf("Added book not found")
	}
}

func TestAddBook_ReturnsErrorIfIDExists(t *testing.T) {
	t.Parallel()

	catalog := getTestCatalog()

	_, ok := catalog.GetBook("ABC04")
	if !ok {
		t.Fatal("Book not present")
	}

	err := catalog.AddBook(books.Book{
		ID:     "ABC04",
		Title:  "Atomic Habit",
		Author: "James Clear",
		Copies: 2,
	})

	if err == nil {
		t.Fatalf("want error for duplicate ID, got nil")
	}
}

func TestSetCopies_OnCatalogModifiesSpecifiedBook(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()

	book, ok := catalog.GetBook("ABC04")
	if !ok {
		t.Fatal("Book not found")
	}

	if book.Copies != 1 {
		t.Fatalf("Want 1 copy before change, got: %d", book.Copies)
	}

	err := catalog.SetCopies("ABC04", 12)
	if err != nil {
		t.Fatal(err)
	}

	book, ok = catalog.GetBook("ABC04")
	if !ok {
		t.Fatal("Book not found")
	}

	if book.Copies != 12 {
		t.Fatalf("Want 12 copy before change, got: %d", book.Copies)
	}

}

// func TestSetCopies_SetsNumberOfCopiesToGivenValue(t *testing.T) {
// 	t.Parallel()
// 	book := books.Book{
// 		Copies: 5,
// 	}

// 	err := book.SetCopies(12)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if book.Copies != 12 {
// 		t.Errorf("want 12 copies, got %d", book.Copies)
// 	}

// }

// func TestSetCopies_ReturnErrorIfCopiesNegative(t *testing.T) {
// 	t.Parallel()

// 	book := books.Book{}

// 	err := book.SetCopies(-1)
// 	if err == nil {
// 		t.Fatalf("Want error for negative copies, got: nil")
// 	}
// }

func assertTestBooks(t *testing.T, got []books.Book) {
	t.Helper()

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

	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})

	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func getTestCatalog() books.Catalog {
	return books.Catalog{
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
