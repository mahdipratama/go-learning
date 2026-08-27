package main

import (
	"books"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: copies <BOOK ID> <NUMBER OF COPIES>")
	}

	catalog, err := books.OpenCatalog("testdata/catalog")
	if err != nil {
		fmt.Printf("Opening catalog: %v \n", err)
	}

	ID := os.Args[1]
	copies, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Invalid number format:", err)
		os.Exit(1)
	}

	// err = catalog.SetCopies(ID, copies)
	// if err != nil {
	// 	fmt.Printf("Updating book: %v \n", err)
	// }

	// err = catalog.Sync("testdata/catalog")
	// if err != nil {
	// 	fmt.Printf("Writing catalog: %v \n")
	// }

	fmt.Printf("Updated book %v to %d copies \n", ID, copies)
	fmt.Printf("%v", catalog)
}
