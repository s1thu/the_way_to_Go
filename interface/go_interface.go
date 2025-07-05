package main

type Speaker interface {
	Speak() string
}

type Dog struct{}

type Cat struct{}

type Human struct{}
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

func main() {
	var s Speaker
	s = Animal{}       // Assigning Animal to Speaker interface
	println(s.Speak()) // Output: Animal speaks

}
