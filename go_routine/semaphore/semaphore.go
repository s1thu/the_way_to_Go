package main

import (
	"fmt"
	"time"
)

func main() {
	sema := make(chan struct{}, 2) // semaphore with 2 slots

	for i := 1; i <= 5; i++ {
		go func(id int) {
			sema <- struct{}{} // acquire
			fmt.Printf("Task %d running\n", id)
			time.Sleep(1 * time.Second)
			fmt.Printf("Task %d done\n", id)
			<-sema // release
		}(i)
	}

	time.Sleep(5 * time.Second)
}
