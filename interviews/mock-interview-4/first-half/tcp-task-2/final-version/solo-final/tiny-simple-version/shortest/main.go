package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	go startServer(":8081")
	time.Sleep(3 * time.Second)
	go startClient("localhost:8081")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if _, err := conn.Write([]byte("1\n")); err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
	}

}

// server
func startServer(port string) {
	netListener, _ := net.Listen("tcp", port)
	defer netListener.Close()

	for {
		connection, _ := netListener.Accept()
		go handleConnection(connection)
	}

}

// client
func startClient(addr string) {
	connection, _ := net.Dial("tcp", addr)
	defer connection.Close()

	buffer := make([]byte, 10)

	for {
		bytesReadNumber, _ := connection.Read(buffer)
		fmt.Print(string(buffer[:bytesReadNumber]))
	}

}
