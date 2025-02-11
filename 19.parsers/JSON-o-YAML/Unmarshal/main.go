package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Интерфейс Parser
type Parser interface {
	Parse(filename string) error
	SaveJSON(filename string) error
	SaveYAML(filename string) error
}

// Структура Post
type Post struct {
	Title       string    `json:"title" yaml:"title"`
	Description string    `json:"description" yaml:"description"`
	Topics      []Topic   `json:"topics" yaml:"topics"`
	Author      Author    `json:"author" yaml:"author"`
	Date        time.Time `json:"date" yaml:"date"`
}

// Структура Author
type Author struct {
	Name     string `json:"name" yaml:"name"`
	Nickname string `json:"nickname" yaml:"nickname"`
}

// Структура Topic
type Topic struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
}

// Реализация метода Parse для структуры Post
func (p *Post) Parse(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Декодирование JSON данных в структуру Post
	err = json.Unmarshal(data, p)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return nil
}

// Реализация метода SaveJSON для структуры Post
func (p *Post) SaveJSON(filename string) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}
	return saveToFile(filename, data)
}

// Реализация метода SaveYAML для структуры Post
func (p *Post) SaveYAML(filename string) error {
	data, err := yaml.Marshal(p)
	if err != nil {
		return fmt.Errorf("error marshalling YAML: %w", err)
	}
	return saveToFile(filename, data)
}

// Вспомогательная функция для сохранения данных в файл
func saveToFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}
	return nil
}

func main() {
	// Создаем новый пост
	post := &Post{}

	// Чтение данных из JSON файла и парсинг
	err := post.Parse("post.json")
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Выводим данные
	fmt.Printf("Post Title: %s\n", post.Title)
	fmt.Printf("Post Description: %s\n", post.Description)
	fmt.Printf("Author: %s (%s)\n", post.Author.Name, post.Author.Nickname)
	fmt.Printf("Date: %s\n", post.Date)

	// Сохранение данных в новый JSON файл
	err = post.SaveJSON("new_post.json")
	if err != nil {
		fmt.Println("Error saving JSON:", err)
		return
	}

	// Сохранение данных в YAML файл
	err = post.SaveYAML("post.yaml")
	if err != nil {
		fmt.Println("Error saving YAML:", err)
		return
	}

	fmt.Println("Data successfully saved to new_post.json and post.yaml")
}
