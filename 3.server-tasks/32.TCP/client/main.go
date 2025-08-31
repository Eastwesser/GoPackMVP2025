package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Подключаемся к TCP серверу
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to TCP server at localhost:8080")
	fmt.Println("Type 'quit' to exit")

	// Горутина для чтения сообщений от сервера
	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Disconnected from server")
				os.Exit(0)
			}
			fmt.Print("Server: " + message)
		}
	}()

	// Читаем ввод пользователя и отправляем на сервер
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		// Отправляем сообщение серверу
		_, err := conn.Write([]byte(text + "\n"))
		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			break
		}

		// Выход если пользователь ввел "quit"
		if strings.ToLower(text) == "exit" {
			fmt.Println("Goodbye!")
			break
		}
	}
}
