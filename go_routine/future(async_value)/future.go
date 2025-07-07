package main

import (
	"fmt"
	"time"
)

func FutureValue() <-chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("Computing in background...")
		time.Sleep(2 * time.Second)
		ch <- 42 // Send result after delay
	}()
	return ch
}

func main() {
	future := FutureValue()
	fmt.Println("Future started, doing other work...")

	// Do something else here...

	result := <-future // Wait for future value
	fmt.Println("Got future result:", result)
}
