package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	file, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parser Result:\nName: %s\nAge: %d\nEmail: %s\n",
		config.Name,
		config.Age,
		config.Email,
	)

}
