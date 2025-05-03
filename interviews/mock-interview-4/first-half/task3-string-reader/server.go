package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ln.Close()

	strings := []string{
		"Успешны те, кто действуют отважно,",
		"Ведь целятся туда, где тяжело.",
		"Знай, честный труд, усердный и прилежный",
		"Пробьется с рвением сверхдалеко.",
	}

	fmt.Println("Waiting for connections...")
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	for _, str := range strings {
		_, err := conn.Write([]byte(str + "\r\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Done.")
}
