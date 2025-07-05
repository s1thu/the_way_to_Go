package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // unbuffered = synchronous

	go func() {
		fmt.Println("Goroutine sending message")
		time.Sleep(2 * time.Second)  // Simulate some work
		ch <- "Hello from goroutine" // blocks until main receives
	}()

	msg := <-ch // blocks until goroutine sends
	fmt.Println(msg)
}
