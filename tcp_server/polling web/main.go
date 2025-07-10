package main

import (
	"fmt"
	"io"
	"net/http"
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

func main() {
	// Executes an HTTP HEAD request for all URL's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}

	res, err := http.Get("http://www.google.com/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response body:", string(data))
}
