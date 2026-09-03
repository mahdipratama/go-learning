package books_test

import (
	"books"
	"cmp"
	"encoding/json"
	"io"
	"net/http"
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

	catalog.Path = t.TempDir() + "/catalog"
	err := catalog.Sync()
	if err != nil {
		t.Fatal(err)
	}

	newCatalog, err := books.OpenCatalog(catalog.Path)
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

func TestSetCopies_SetsNumberOfCopiesToGivenValue(t *testing.T) {
	t.Parallel()
	book := books.Book{
		Copies: 5,
	}

	err := book.SetCopies(12)
	if err != nil {
		t.Fatal(err)
	}

	if book.Copies != 12 {
		t.Errorf("want 12 copies, got %d", book.Copies)
	}

}

func TestSetCopies_ReturnErrorIfCopiesNegative(t *testing.T) {
	t.Parallel()

	book := books.Book{}

	err := book.SetCopies(-1)
	if err == nil {
		t.Fatalf("Want error for negative copies, got: nil")
	}
}

func TestSetCopies_IsRaceFree(t *testing.T) {
	t.Parallel()

	catalog := getTestCatalog()

	go func() {
		for range 100 {
			err := catalog.SetCopies("ABC04", 0)
			if err != nil {
				panic(err)
			}
		}
	}()

	for range 100 {
		_, err := catalog.GetCopies("ABC04")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestNewCatalog_GetEmptyCatalog(t *testing.T) {
	t.Parallel()
	catalog := books.NewCatalog()
	books := catalog.GetAllBooks()
	if len(books) > 0 {
		t.Errorf("want empty catalog, got %v", books)
	}

}

func TestServer_ListAllBooks(t *testing.T) {
	catalog := getTestCatalog()
	catalog.Path = t.TempDir() + "/catalog"

	go func() {
		err := books.ListenAndServe(":3000", catalog)
		if err != nil {
			panic(err)
		}
	}()

	resp, err := http.Get("http://localhost:3000/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status %d", resp.StatusCode)
	}

	bookList := []books.Book{}

	// err = json.NewDecoder(resp.Body).Decode(&bookList)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(data, &bookList)
	if err != nil {
		t.Fatalf("%v in %q", err, data)
	}

	assertTestBooks(t, bookList)
}

func getTestCatalog() *books.Catalog {
	catalog := books.NewCatalog()

	err := catalog.AddBook(books.Book{
		Title:  "Never Finished",
		Author: "David Goggins",
		Copies: 2,
		ID:     "ABC03",
	})

	if err != nil {
		panic(err)
	}

	err = catalog.AddBook(books.Book{
		Title:  "The Mountain is You",
		Author: "Briana Weist",
		Copies: 1,
		ID:     "ABC04",
	})

	if err != nil {
		panic(err)
	}

	return catalog
}

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
