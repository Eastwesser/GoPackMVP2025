package main

import "fmt"

//Композиция позволяет создавать сложные объекты, комбинируя более простые.
//Например, ученик может иметь список навыков и учебных материалов.

// Навык
type Skill struct {
	Name        string
	Description string
}

// Учебный материал
type Material struct {
	Title string
	Type  string
}

// Ученик
type Student struct {
	Name      string
	Skills    []Skill
	Materials []Material
}

func main() {
	student := Student{
		Name: "Иван Иванов",
		Skills: []Skill{
			{Name: "Speaking", Description: "Разговорный английский"},
			{Name: "Reading", Description: "Чтение текстов"},
		},
		Materials: []Material{
			{Title: "Basic English Grammar", Type: "Книга"},
			{Title: "Daily Conversations", Type: "Аудио"},
		},
	}

	fmt.Println("Ученик:", student.Name)
	fmt.Println("Навыки:")
	for _, skill := range student.Skills {
		fmt.Println("-", skill.Name, ":", skill.Description)
	}
	fmt.Println("Материалы:")
	for _, material := range student.Materials {
		fmt.Println("-", material.Title, "(", material.Type, ")")
	}
}
