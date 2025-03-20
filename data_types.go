package main

import "fmt"

type GenderStatus string

const (
	Male GenderStatus = "Male"
	Female = "Female"
)

type Person struct{
	Name string
	age int
	Gender GenderStatus
}

func (g GenderStatus) getGenderStatus() string{
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown"
	}
}

// Function to demonstrate elementary types
func elementary_types() {
	// Boolean
	var isGoFun bool = true
	fmt.Println("Boolean:", isGoFun) // Output: Boolean: true

	// Integer
	var age int = 25
	fmt.Println("Integer:", age) // Output: Integer: 25

	// Floating Point
	var price float64 = 99.99
	fmt.Println("Float:", price) // Output: Float: 99.99

	// String
	var name string = "Go Language"
	fmt.Println("String:", name) // Output: String: Go Language

	// Complex Number
	var complexNum complex128 = complex(3, 4)
	fmt.Println("Complex Number:", complexNum) // Output: (3+4i)
}

func structure_types(){
	Person1 := Person{"Sithu",26,Male};
	fmt.Println("Person1:",Person1)

	// Array
	var numbers [3]int = [3]int{10, 20, 30}
	fmt.Println(numbers)      // Output: [10 20 30]
	fmt.Println(numbers[0])   // Output: 10

	// Slice
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)       // Output: [Red Green Blue]

	// Map
	var person = map[string]string{
		"name": "Sithu",
		"age": "26",
	}
	fmt.Println(person)       // Output: map[age:26 name:Sithu]
	fmt.Println(person["name"]) // Output: Sithu

	// Channel
	var channel = make(chan string)
	fmt.Println(channel)      // Output: 0xc0000b6000

	go func(){
		channel <- "Hello"
	}()

	message := <- channel
	fmt.Println(message)      // Output: Hello
}