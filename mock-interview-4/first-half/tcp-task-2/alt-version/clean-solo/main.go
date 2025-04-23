package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	go runServer()
	time.Sleep(100 * time.Millisecond)
	runClient()
}

func runServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Server error:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := conn.Write([]byte("2\n"))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func runClient() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Client error:", err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		fmt.Print("Client received:", msg)
	}
}
