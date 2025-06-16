package main

import "fmt"

func reslicing(){

	var arr = [15]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice1 = arr[5:10] // Slice from index 5 to 9
	fmt.Println(slice1)
	 slice1 = slice1[1:3] // Resl
	 fmt.Println(slice1) // Output: [6 7]
}