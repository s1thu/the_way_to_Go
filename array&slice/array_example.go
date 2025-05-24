package main

import "fmt"

func array(){
	var a = new([5]int) //declaration of an array
	a[0] = 1 //initialization of an array
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	//Array literals and parameters
	//1st Variant
	b := [5]int{1, 2, 3, 4, 5} //declaration of an array
	//2nd Variant
	c := [...]int{1, 2, 3, 4, 5} //declaration of an array
	//3rd Variant
	d := [5]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5} //declaration of an array
	
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	
}