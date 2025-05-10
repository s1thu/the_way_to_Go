package main

import "fmt"

func pointer_test() {
	s := "HelloWorld"
	fmt.Println(s)
	var p *string = &s
	fmt.Println(*p)
	*p = "SithuWin"
	fmt.Println(*p)
}
