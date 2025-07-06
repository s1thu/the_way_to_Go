package main

import (
	"fmt"
	"time"
)

func main1() {

	time.Sleep(3 * time.Second)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
	}()

	go func() {
		defer close(ch2)
		ch2 <- 2
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 closed")
				ch1 = nil // disable this case
			} else {
				fmt.Println("from ch1:", v)
			}
		case v, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 closed")
				ch2 = nil
			} else {
				fmt.Println("from ch2:", v)
			}
		}
	}
}
