package main

import (
	"fmt"
	"time"
)

func goroutineB() {
	for i := range 10 {
		fmt.Println("Hello from Go routines B!", i)
		time.Sleep(10 * time.Millisecond)

	}
}

func main() {
	go goroutineB()
	for i := range 10 {
		fmt.Println("Hello from Go routines A!", i)
		time.Sleep(10 * time.Millisecond)
	}
}
