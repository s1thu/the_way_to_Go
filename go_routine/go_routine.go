package main

import (
	"fmt"
	"sync"
	"time"
)

func longWait(wg *sync.WaitGroup) {
	fmt.Println("Long wait started")
	time.Sleep(5 * time.Second) // Simulate a long wait
	fmt.Println("Long wait finished")
	defer wg.Done() // Notify that this goroutine is done
}

func shortWait(wg *sync.WaitGroup) {
	fmt.Println("Short wait started")
	time.Sleep(2 * time.Second) // Simulate a short wait
	fmt.Println("Short wait finished")
	defer wg.Done() // Notify that this goroutine is done
}

func main() {
	fmt.Println("In main()")
	wg := &sync.WaitGroup{} // Create a new WaitGroup
	wg.Add(2)               // Add two goroutines to the wait group
	go longWait(wg)
	go shortWait(wg)
	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("At the end of main()")
}
