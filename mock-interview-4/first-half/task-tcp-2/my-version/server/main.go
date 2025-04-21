package main

import (
	"fmt"
	"net"
	"time"
)

func handleCon(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := conn.Write([]byte("1\n"))
		if err != nil {
			fmt.Println("error of writing", err)
			return
		}
		time.Sleep(1 * time.Second)
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error listening", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on port :8081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accepting error", err)
			continue
		}
		go handleCon(conn)
	}

}
