package main

import "fmt"

func test_function() {
	one()
}

func one() {
	fmt.Println("Hello")
	defer two()
	fmt.Println("One")
}
func two() {
	fmt.Println("Hello")
	fmt.Println("Two")
}
