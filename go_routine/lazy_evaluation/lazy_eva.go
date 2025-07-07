package main

import "fmt"

// LazySum returns a function that calculates sum only when called
func LazySum(a, b int) func() int {
	return func() int {
		fmt.Println("Evaluating sum now...")
		return a + b
	}
}

func main() {
	sum := LazySum(5, 7) // Not yet evaluated
	fmt.Println("Sum function created, not evaluated.")

	result := sum() // Evaluated here
	fmt.Println("Result:", result)
}
