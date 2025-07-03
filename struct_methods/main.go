package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

type T struct {
	a int `info:This is a tag`
	b int `info:A raw string tag`
	c int `key1:"value1" key2:"value2"`
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) GetName() string {
	return p.Name
}

type TagType struct { // tags
	field1 bool   `info:"An important answer"`
	field2 string `info:"The name of the thing"`
	field3 int    `info:"How much there are"`
}

type Anonymous struct {
	int
	string
}

func main() {
	//a factory of a struct
	p := NewPerson("Alice", 30)
	fmt.Println("Name:", p.Name)
	println("Age:", p.Age)

	tt := TagType{true, "Barack Obama", 1}
	for i := range 3 {
		refTag(tt, i)
	}

	t := T{}
	fmt.Println(reflect.TypeOf(t).Field(0).Tag)
	if field, ok := reflect.TypeOf(t).FieldByName("b"); ok {
		fmt.Println(field.Tag)
	}
	if field, ok := reflect.TypeOf(t).FieldByName("c"); ok {
		fmt.Println(field.Tag)
		fmt.Println(field.Tag.Get("key1"))
	}
	if field, ok := reflect.TypeOf(t).FieldByName("d"); ok {
		fmt.Println(field.Tag)
	} else {
		fmt.Println("Field not found")
	}

	a := Anonymous{int: 42, string: "Hello"}
	fmt.Println("Anonymous struct:", a)
	fmt.Println("Anonymous struct int field:", a.int)
	fmt.Println("Anonymous struct string field:", a.string)

	personone := Person{Name: "John", Age: 25}
	persontwo := Person{Name: "Jane", Age: 30}
	fmt.Println("Person One:", personone.GetName()) //method in go lang
	fmt.Println("Person Two:", persontwo.GetName())
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)     // getting field at a position ix
	fmt.Printf("%v\n", ixField.Tag) // printing tags
}
