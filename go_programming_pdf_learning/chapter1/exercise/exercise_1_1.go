package main

import (
	"fmt"
	"os"
	"strings"
)

func exercise_1_1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
