package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Marshal вместо Unmarshal - преобразует структуры Go в байты

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	person := Person{
		Name:  "Иван Иванов",
		Age:   30,
		Email: "ivan@example.com",
	}

	// Кодируем в JSON с отступами
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Записываем в файл
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Данные успешно записаны в output.json")
	fmt.Println(string(jsonData))
}
