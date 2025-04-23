package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	go func() {
		netListener, _ := net.Listen("tcp", ":8081") // Запуск сервера
		for {
			connection, _ := netListener.Accept()

			go func(c net.Conn) {
				defer c.Close()
				for {
					c.Write([]byte("1\n"))
					time.Sleep(time.Second)
				}
			}(connection)

		}
	}()

	time.Sleep(3 * time.Second) // Даем серверу время на запуск

	connection, _ := net.Dial("tcp", "localhost:8081") // Запуск клиента
	defer connection.Close()
	buffer := make([]byte, 10)
	for {
		bytesReadNumber, _ := connection.Read(buffer)
		fmt.Print(string(buffer[:bytesReadNumber]))
	}
}
