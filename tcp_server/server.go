package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the TCP server...")

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return // terminate program
		}
		fmt.Printf("Received data: %v", string(buf))
	}
}
