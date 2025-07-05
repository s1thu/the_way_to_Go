package main

import (
	"fmt"
)

type flt func(int) bool

// isOdd takes an int slice and returns a bool set to true if the
// int parameter is odd, or false if not.
// isOdd is of type func(int) bool which is what flt is declared to be.

func isOdd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true
}

// Same comment for isEven
func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

// higher-order function
func calculate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func filter(sl []int, f flt) []int {
	var res []int
	for _, val := range sl {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

type Car struct {
	Model string
	Year  int
}

type Cars []*Car

func (c *Car) Process(f func(*Car)) {
	f(c)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)

	sum := calculate(5, 10, func(a, b int) int {
		return a + b
	})
	fmt.Println("Sum of 5 and 10 is: ", sum)

	carOne := Car{Model: "Toyota", Year: 2020}
	carOne.Process(func(c *Car) {
		fmt.Printf("Car Model: %s, Year: %d\n", c.Model, c.Year)
	})

}
