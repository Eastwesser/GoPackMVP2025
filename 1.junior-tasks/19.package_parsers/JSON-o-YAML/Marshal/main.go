package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Marshal вместо Unmarshal - преобразует структуры Go в байты

type Person struct {
	Name  string `json:"name" yaml:"name"`
	Age   int    `json:"age" yaml:"age"`
	Email string `json:"email" yaml:"email"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите формат: json или yaml")
	}

	format := strings.ToLower(os.Args[1])
	person := Person{
		Name:  "Алексей Сидоров",
		Age:   35,
		Email: "alex@example.com",
	}

	var data []byte
	var err error
	outputFile := ""

	switch format {
	case "json":
		data, err = json.MarshalIndent(person, "", "  ")
		outputFile = "output.json"
	case "yaml":
		data, err = yaml.Marshal(person)
		outputFile = "output.yaml"
	default:
		log.Fatal("Неподдерживаемый формат. Используйте json или yaml")
	}

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(outputFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Данные успешно записаны в %s\n", outputFile)
	fmt.Println(string(data))
}
