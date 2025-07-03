package main

type Person struct{
	Name string
	Age int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age: age,
	}
}
func main(){
	//a factory of a struct
	p := NewPerson("Alice", 30)
	println("Name:", p.Name)
	println("Age:", p.Age)
}