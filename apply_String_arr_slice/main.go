package main

import "fmt"

func main(){
	//Making slice of bytes or runes from a string
	var s = "Hello, World!"
	c := []byte(s)
	for _, values := range c {
		//Printing the byte values
		println(values)
	}

	//with copy
	fmt.Println("\nUsing copy to create a new slice from the original byte slice:")
	d := make([]byte, len(c))
	copy(d, s)
	for _, values := range d {
		//Printing the byte values
		println(values)
	}
}