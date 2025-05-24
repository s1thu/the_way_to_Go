package main

import "fmt"

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

	fmt.Print(task(7,9,3,5,1))
	printrec(1)
}

func n() {
	print(a)
}
func m() {
	a = "O"
	print(a)
}

func task(a...int) int {
    if len(a) == 0 {
        return 0
    }
    value := a[0]
    for _, v := range a {
		fmt.Println(v)
        if v < value {
            value = v
        }
    }
    return value
}

func printrec(i int) {
    if i > 10 {
        return
    }
    printrec(i + 1)
    fmt.Printf("%d ", i)
}