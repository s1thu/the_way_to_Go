package main

import (
	"bytes"
	"fmt"
)

func byte_learn(){
	var b bytes.Buffer
	b.WriteString("Hello, World!")

	fmt.Fprintf(&b, " %d bytes written", b.Len())
	fmt.Println(b.String())
}