package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	
	listener, err := net.Listen("tcp", "<<USE YOUR IP>>:<PORT>")
	if err != nil {
        fmt.Println("Error:", err)
        return
    }
	defer listener.Close()

	fmt.Println("Server is listening on IP(local) = <<USE YOUR IP>> and port = <PORT>")

	for {
        // Accept incoming connections
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        // Handle client connection in a goroutine
        go handleClient(conn)
    }
}

func elapsedTime(start *time.Time) {
	et := time.Since(*start)
	fmt.Println("Server Time Taken = ", et)
}

func handleClient(conn net.Conn) {
    defer conn.Close()

	// Create a buffer to read data into
 	buffer := make([]byte, 1024)

	start := time.Now()
	defer elapsedTime(&start)

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		// Process and use the data (here, we'll just print it)
		fmt.Printf("%s", buffer[:n])
	}
}
