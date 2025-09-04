package main

import (
	"bufio"
	"fmt"
	"net"
)

func TCPServer(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing TCP connection")
		}
	}(conn)
	fmt.Printf("Client connected to %s\n", conn.RemoteAddr().String())

	message := fmt.Sprintf(
		"Hello, %s!\n",
		conn.RemoteAddr().String(),
	)
	_, err := conn.Write([]byte(message))
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cliMessage := scanner.Text()
		fmt.Printf("%s: %s\n", conn.RemoteAddr().String(), cliMessage)

		response := fmt.Sprintf("Echo: %s\n", cliMessage)
		_, err := conn.Write([]byte(response))
		if err != nil {
			return
		}

		if cliMessage == "exit" {
			fmt.Printf("Client disconnected: %s\n", conn.RemoteAddr().String())
			break
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error listening on port 8080: %s\n", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)

	fmt.Printf("Listening on port 8080\n")
	fmt.Println("waiting for client to connect...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}
		go TCPServer(conn)
	}
}
