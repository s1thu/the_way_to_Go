package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "Dog barks"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Cat meows"
}

type Human struct{}

func (h *Human) Speak() string {
	return "Human speaks"
}

type Animal struct {
	Dog
	Cat
}

type NotAnimal struct {
	Human
}

func (a NotAnimal) Speak() string {
	return "NotAnimal speaks"
}

func (a Animal) Speak() string {
	return "Animal speaks"
}

// empty interface
func printAnything(i interface{}) {
	fmt.Println(i)
}

func main() {
	var s Speaker

	s = &Dog{}         // Assigning Animal to Speaker interface
	println(s.Speak()) // Output: Animal speaks

	printAnything("hello")
	printAnything(1)

	//type asssertion
	var i interface{} = "Go"
	str, ok := i.(string)
	if ok {
		fmt.Println(str) // Go
	}

}
