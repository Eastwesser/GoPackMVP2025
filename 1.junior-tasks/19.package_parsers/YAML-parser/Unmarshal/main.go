package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name  string `yaml:"name"`
	Age   int    `yaml:"age"`
	Email string `yaml:"email"`
}

func main() {
	file, err := os.ReadFile("data.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("YAML Parser Result:\nName: %s\nAge: %d\nEmail: %s\n",
		config.Name,
		config.Age,
		config.Email,
	)

}
