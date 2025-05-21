package main

import (
	"fmt"
	"log"
	"os"
)

// trace logs entering and leaving of a function.
func trace(msg string) func() {
	log.Println("enter:", msg)
	return func() {
		log.Println("leave:", msg)
	}
}

// openFile opens a file and returns a file pointer.
func openFile(filename string) (*os.File, error) {
	defer trace("openFile")()
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// processFile processes the opened file.
func processFile(file *os.File) error {
	defer trace("processFile")()
	content, err := os.ReadFile(file.Name())
	if err != nil {
		return err
	}
	fmt.Println("File content:\n", string(content))
	return nil
}

// closeFile closes the file.
func closeFile(file *os.File) {
	defer trace("closeFile")()
	if file != nil {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}
}

func defer_tracing() {
	defer trace("main")()

	fileName := "function.go"

	file, err := openFile(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	// Ensure the file is closed no matter what.
	defer closeFile(file)

	err = processFile(file)
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}

	fmt.Println("File processed successfully.")
}
