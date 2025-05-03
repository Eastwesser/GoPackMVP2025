package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server", err)
		os.Exit(1)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("String: %s, Length: %d\n", line, len(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading data:", err)
	}

}
