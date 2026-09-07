package main

import (
	"books"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// os.Args is a slice of strings, came from user input as an arguments then store it
	// os.Args[0]: Contains the name or path of the running program executable.
	// os.Args[1:]: Contains all the actual user-provided arguments passed to the program.

	// HOW TO USE:
	// Current Folder -> go run main.go [bookId]
	// root -> go run ./cmd/find [bookId]

	if len(os.Args) != 2 {
		fmt.Println("Usage: find <BOOK ID>")
		return
	}

	ID := os.Args[1]

	resp, err := http.Get("http://localhost:3000/v1/find/" + ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status %d", resp.StatusCode)
		return
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := books.Book{}

	err = json.Unmarshal(data, &book)
	if err != nil {
		fmt.Printf("%v in %q ", err, data)
		return
	}

	fmt.Println(book)

}
