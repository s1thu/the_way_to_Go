package main

import (
	"fmt"
	"os"
)

func exercise_1_2() {

	for index, v := range os.Args[0:] {
		fmt.Println("Index:", index, "Arg:", v)
	}
}
