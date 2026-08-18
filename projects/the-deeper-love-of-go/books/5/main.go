package main

import "fmt"

func main() {
	var title = "The Mountain is You"
	var author = "Briana Weist"

	fmt.Println("Books in stock")
	printBook(title, author)

	title = "Can't Hurt me"
	author = "David Goggins"
	printBook(title, author)
}

func printBook(title, author string) {

	fmt.Println(title, "by", author)

}
