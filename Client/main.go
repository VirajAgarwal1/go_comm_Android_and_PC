package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func Hello() string {
	return "Hello, Android World\n\t-Golang"
}

func send_sq_pc(conn *net.Conn, x int, wg *sync.WaitGroup) error {

	// Send data to the server
	defer wg.Done()
	data := []byte("Hello, from the Android to Server! Here is perfect square = " + strconv.Itoa(x*x) + "\n")

	_, err := (*conn).Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func main() {
	conn, err := net.Dial("tcp", "<<USE YOUR IP>>:<PORT>")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go send_sq_pc(&conn, i, &wg)
	}
	wg.Wait()

	exec_time := time.Since(start)
	fmt.Println("Local Time Taken = ", exec_time)

	fmt.Println(Hello())
}
