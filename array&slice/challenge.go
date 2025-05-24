package main

import "fmt"

func challenge() {
    var arr [15]int // Declare an array of size 15

    // Fill the array using a loop
    for i := 0; i < len(arr); i++ {
        arr[i] = i
    }

    // Print the filled array
    fmt.Println(arr)
}
