package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprint(w, "Hello, "+req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("0.00.0:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
