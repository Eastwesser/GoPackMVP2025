package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	go func() {
		serverListener, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
		defer serverListener.Close()
		log.Println("Server is listening port :8081")

		for {
			conn, err := serverListener.Accept()
			if err != nil {
				log.Printf("Accept error: %v", err)
				continue
			}
			go handleServerConnection(conn)
		}
	}()

	// Даем серверу время на запуск
	time.Sleep(5 * time.Second)

	// Запускаем клиент
	conn, err := net.Dial("tcp", "localhost:8081") // _ -> err
	if err != nil {
		log.Fatalf("Client failed to connect: %v", err)
	}
	defer conn.Close()
	log.Println("Client connected to server")

	// Горутина для чтения сообщений от сервера
	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Client read error: %v", err)
				return
			}
			//fmt.Printf("Received: %s", msg)
			fmt.Printf("%s", msg)
		}
	}()

	// Основной поток - можно добавить ввод с клавиатуры
	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func handleServerConnection(conn net.Conn) {
	defer conn.Close()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if _, err := conn.Write([]byte("1\n")); err != nil {
			log.Printf("Server write error: %v", err)
			return
		}
	}
}
