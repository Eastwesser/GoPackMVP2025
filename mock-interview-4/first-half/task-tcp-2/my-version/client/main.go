package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		fmt.Println("Failed to connect", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error of reading", err)
			return
		}
		fmt.Print("Received ", message)
	}

}
