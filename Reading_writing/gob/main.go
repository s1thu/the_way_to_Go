package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Define a sample struct
type User struct {
	Name string
	Age  int
}

func main() {
	// Original data
	original := User{Name: "Alice", Age: 30}

	// Create a buffer to hold the encoded data
	var buf bytes.Buffer

	// Create a gob encoder writing to the buffer
	enc := gob.NewEncoder(&buf)

	// Encode the struct into the buffer
	err := enc.Encode(original)
	if err != nil {
		fmt.Println("Encoding failed:", err)
		return
	}

	fmt.Println("✅ Serialized data into buffer.")

	// Create a decoder reading from the same buffer
	dec := gob.NewDecoder(&buf)

	// Create a variable to hold the decoded data
	var decoded User

	// Decode into it
	err = dec.Decode(&decoded)
	if err != nil {
		fmt.Println("Decoding failed:", err)
		return
	}

	fmt.Println("✅ Decoded data from buffer:")
	fmt.Println("Name:", decoded.Name)
	fmt.Println("Age:", decoded.Age)
}
