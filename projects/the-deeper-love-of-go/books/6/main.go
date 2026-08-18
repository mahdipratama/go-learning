package main

import "fmt"

func main() {
	var title = "The Mountain is You"
	var author = "Briana Weist"
	var copies = 1

	fmt.Println("Books in stock")
	printBook(title, author, copies)

	title = "Can't Hurt me"
	author = "David Goggins"
	copies = 2
	printBook(title, author, copies)
}

func printBook(title, author string, copies int) {
	fmt.Println(title, "by", author, "-", copies, "copies")

}
