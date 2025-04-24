package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	go startServer(":8081")          // Запуск сервера
	time.Sleep(3 * time.Second)      // Даем серверу время на запуск
	go startClient("localhost:8081") // Запуск клиента
	// Ожидание сигнала завершения
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
			fmt.Printf("error: %v\n", err) // error while writing
			return
		}
	}

}

// Серверная часть
func startServer(port string) {
	netListener, _ := net.Listen("tcp", port) // _ -> err
	//if err != nil {
	//	fmt.Printf("Server failed to start: %v\n", err)
	//	return
	//}
	defer netListener.Close()

	for {
		connection, _ := netListener.Accept() // _ -> err
		//if err != nil {
		//	fmt.Printf("Accept error: %v\n", err)
		//	continue
		//}
		go handleConnection(connection)
	}
}

// Клиентская часть
func startClient(addr string) {
	connection, _ := net.Dial("tcp", addr) // _ -> err
	//if err != nil {
	//	fmt.Printf("Client failed to connect: %v\n", err)
	//	return
	//}
	defer connection.Close()

	buffer := make([]byte, 10)

	for {
		bytesReadNumber, _ := connection.Read(buffer) // _ -> err
		//if err != nil {
		//	fmt.Printf("Read error: %v\n", err)
		//	return
		//}
		fmt.Print(string(buffer[:bytesReadNumber]))
	}

}
