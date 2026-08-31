package main

import "fmt"

func goroutineB() {
	for i := range 10 {
		fmt.Println("Hello from Go routines B!", i)
	}
}

func main() {
	go goroutineB()
	for i := range 10 {
		fmt.Println("Hello from Go routines A!", i)
	}
}
