package main

import "fmt"

// aliasing type
type Celsius float32
type Fahrenheit float32

// Function to convert celsius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {
	
	var temp Fahrenheit	
	temp = Fahrenheit((t * 9 / 5) + 32)
	fmt.Printf("%f Celsius = %f Fahrenheit\n", t, temp)
	return temp

}