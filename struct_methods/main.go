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

type TagType struct { // tags
	field1 bool   `info:"An important answer"`
	field2 string `info:"The name of the thing"`
	field3 int    `info:"How much there are"`
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
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)     // getting field at a position ix
	fmt.Printf("%v\n", ixField.Tag) // printing tags
}
