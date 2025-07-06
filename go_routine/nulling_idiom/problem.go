package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
	}()

	for {
		select {
		// case v, ok := <-ch1:
		// 	if !ok {
		// 		fmt.Println("channel closed")
		// 		return
		// 	}
		// 	fmt.Println("from ch1:", v)
		case v := <-ch1:
			fmt.Println(v)
		}
	}
	// time.Sleep(3 * time.Second)
}
