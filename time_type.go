package main

import (
	"fmt"
	"time"
)

func time_test() {
	t := time.Now().UTC()
	fmt.Print(t.UTC()) // e.g.: 29.10.2019
}
