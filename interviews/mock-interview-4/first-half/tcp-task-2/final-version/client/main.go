package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const serverAddr = "localhost:8080"

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}
		fmt.Printf("Received: %s", msg)
	}
}
