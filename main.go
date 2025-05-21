package main

var a = "G"

func main() {
	// elementary_types() // Calling the function from data_types.go
	// structure_types()
	//toFahrenheit(100.0) // Example usage of the toFahrenheit function;
	//string_example()
	strings_functions()
	time_test()
	pointer_test()
	n()

	m()
	n()
	test_function()
	defer_tracing()
}

func n() {
	print(a)
}
func m() {
	a = "O"
	print(a)
}
