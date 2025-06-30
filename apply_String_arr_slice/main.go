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

		sla := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	fmt.Println("before sort: ",sla)
	// sla is passed via call by value, but since sla is a reference type
	// the underlying array is changed (sorted)
	bubbleSort(sla)
	fmt.Println("after sort:  ",sla)

	// reverse a string:
	str := "Google"

	fmt.Printf("The reversed string using variant 1 is -%s-\n", reverse1(str))
	fmt.Printf("The reversed string using variant 2 is -%s-\n", reverse2(str))
	fmt.Printf("The reversed string using variant 3 is -%s-\n", reverse3(str))
}