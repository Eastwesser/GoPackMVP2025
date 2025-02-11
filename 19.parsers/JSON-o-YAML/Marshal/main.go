//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"os"
//	"time"
//)
//
//type Parser interface {
//	Parse(filename string)
//	SaveJSON(filename string) error
//	SaveYAML(filename string) error
//}
//
//type Post struct {
//	Title       string    `json:"title" yaml:"title"`
//	Description string    `json:"description" yaml:"description"`
//	Topics      []Topic   `json:"topics" yaml:"topics"`
//	Author      Author    `json:"author" yaml:"author"`
//	Date        time.Time `json:"date" yaml:"date"`
//}
//
//type Author struct {
//	Name     string `json:"name" yaml:"name"`
//	Nickname string `json:"nickname" yaml:"nickname"`
//}
//
//type Topic struct {
//	Title       string `json:"title" yaml:"title"`
//	Description string `json:"description" yaml:"description"`
//}
//
//func (p *Post) Parse(filename string) {
//	// Пример парсинга данных в структуру Post
//	p.Title = "Go Interface Example"
//	p.Description = "This is a blog post example."
//	p.Topics = []Topic{{Title: "Introduction", Description: "Introduction to Go Interfaces"}}
//	p.Author = Author{Name: "John Doe", Nickname: "JD"}
//	p.Date = time.Now()
//}
//
//func (p *Post) SaveJSON(filename string) error {
//	data, err := json.MarshalIndent(p, "", "  ")
//	if err != nil {
//		return fmt.Errorf("Error marshalling JSON: %w", err)
//	}
//	return saveToFile(filename, data)
//}
//
//func (p *Post) SaveYAML(filename string) error {
//	data, err := yaml.Marshal(p)
//	if err != nil {
//		return fmt.Errorf("Error marshalling YAML: %w", err)
//	}
//	return saveToFile(filename, data)
//}
//
//func saveToFile(filename string, data []byte) error {
//	err := os.WriteFile(filename, data, 0644)
//	if err != nil {
//		fmt.Println("Error writing file:", err)
//		return err
//	}
//	return nil
//}
//
//func main() {
//	// Пример использования
//	post := &Post{}
//	post.Parse("example.json")
//
//	// Сохранение данных в JSON
//	err := post.SaveJSON("post.json")
//	if err != nil {
//		fmt.Println("Error saving JSON:", err)
//	}
//
//	// Сохранение данных в YAML
//	err = post.SaveYAML("post.yaml")
//	if err != nil {
//		fmt.Println("Error saving YAML:", err)
//	}
//}
