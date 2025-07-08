package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//using Scanln
	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello", name)

	//using buffered reader
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age:")
	intput, _ := inputReader.ReadString('\n')
	fmt.Println("Your age is", intput)

	//reading from a file
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +

			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	buff, err := os.ReadFile("input.dat")
	if err != nil {
		fmt.Printf("An error occurred on reading the inputfile\n")
	}
	fmt.Printf("The content of the file is: %s\n", string(buff))
	defer inputFile.Close()

	//reading from column
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(inputFile, &v1, &v2, &v3) // scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		fmt.Printf("The input was: %s", inputString)
	}
}
