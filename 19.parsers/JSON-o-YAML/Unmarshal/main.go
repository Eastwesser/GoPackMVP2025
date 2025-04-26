package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name  string `json:"name" yaml:"name"`
	Age   int    `json:"age" yaml:"age"`
	Email string `json:"email" yaml:"email"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide file path")
	}

	filePath := os.Args[1]
	ext := strings.ToLower(filepath.Ext(filePath))

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	switch ext {
	case ".json":
		if err := json.Unmarshal(file, &config); err != nil {
			log.Fatal(err)
		}
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(file, &config); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unsupported file format")
	}

	fmt.Printf("Universal Parser Result:\nName: %s\nAge: %d\nEmail: %s\n",
		config.Name,
		config.Age,
		config.Email,
	)

}
