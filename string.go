package main

import (
	"fmt"
	"strings"
)

func string_example() {
	var name string = "Hello World"
	fmt.Println(name)

	var doneExt string = "one.done"
	var substr string = "done"
	fmt.Println(strings.HasSuffix(doneExt, ".done"))
	fmt.Println(strings.HasPrefix(doneExt, "one"))
	fmt.Println(strings.Contains(doneExt, substr))

}
