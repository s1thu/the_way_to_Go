package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from channel 2"
	}()

	// Using select to wait for either channel to be ready
	// If both channels are ready, it will select one at random
	// If neither channel is ready, it will go to the default case
	// This is a blocking operation, so it will wait until one of the channels is ready
	// If you want to avoid blocking, you can use a default case
	// Here, we use a default case to avoid blocking if neither channel is ready
	// Note: The select statement will block until one of the channels is ready
	// If you want to handle multiple channels, you can use a for loop with select
	// This will allow you to handle multiple channels in a non-blocking way
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	default:
		fmt.Println("No channels ready")
	}
}
