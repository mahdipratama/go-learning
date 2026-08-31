package main

import (
	"fmt"
	"time"
)

var message = "Hello"

func goroutineB() {
	message = "Goodbye"

	for range 10 {
		fmt.Println(message, "from Go routines B!")
		time.Sleep(10 * time.Millisecond)

	}
}

func main() {
	go goroutineB()
	for range 10 {
		fmt.Println(message, "from Go routines A!")
		time.Sleep(10 * time.Millisecond)
	}
}
