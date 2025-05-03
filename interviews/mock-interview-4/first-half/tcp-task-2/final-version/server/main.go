package main

import (
	"log"
	"net"
	"time"
)

const (
	port    = ":8080"
	message = "2\n"
	delay   = 1 * time.Second
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for range ticker.C {
		_, err := conn.Write([]byte(message))
		if err != nil {
			log.Printf("Write error: %v", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	log.Printf("Server started on %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
