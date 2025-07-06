package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	select {
	case v := <-ch1:
		fmt.Println(v)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: No data received from channel")
	}

	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()
	for {
		select {
		case v := <-ch1:
			fmt.Println("Received:", v)
		case <-tick.C:
			fmt.Println("Tick: No data received from channel", <-tick.C)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout: No data received for 5 seconds, exiting")
			return
		}
	}
}
