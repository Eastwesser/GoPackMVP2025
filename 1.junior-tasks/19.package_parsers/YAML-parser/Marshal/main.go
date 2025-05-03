package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Marshal вместо Unmarshal - преобразует структуры Go в байты

type Person struct {
	Name  string `yaml:"name"`
	Age   int    `yaml:"age"`
	Email string `yaml:"email"`
}

func main() {
	person := Person{
		Name:  "Мария Петрова",
		Age:   25,
		Email: "maria@example.com",
	}

	// Кодируем в YAML
	yamlData, err := yaml.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	// Записываем в файл
	err = os.WriteFile("output.yaml", yamlData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Данные успешно записаны в output.yaml")
	fmt.Println(string(yamlData))
}
