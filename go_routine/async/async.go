package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2) // buffered = async

	go func() {
		ch <- "Hello from goroutine One"
	}()

	go func() {
		ch <- "Hello from goroutine Two"
	}()

	go func() {
		time.Sleep(10 * time.Second) // Simulate some work
		ch <- "Hello from goroutine Three"
	}()
	// ch <- "Third" // would block here if buffer full

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
