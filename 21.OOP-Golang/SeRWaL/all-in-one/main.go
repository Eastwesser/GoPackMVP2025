package main

import "fmt"

// Абстракция: интерфейс для учеников
type Practiceable interface {
	Practice()
}

// Базовый класс (Инкапсуляция)
type BaseStudent struct {
	Name     string
	progress int
}

func (b BaseStudent) GetProgress() int {
	return b.progress
}

func (b *BaseStudent) UpdateProgress(amount int) {
	if amount > 0 {
		b.progress += amount
		fmt.Println(b.Name, "улучшил прогресс на", amount, "%")
	}
}

// Композиция: ученик с навыками и материалами
type Skill struct {
	Name        string
	Description string
}

type Material struct {
	Title string
	Type  string
}

// Наследование (встраивание) и полиморфизм
type AdvancedStudent struct {
	BaseStudent
	Skills    []Skill
	Materials []Material
	Focus     string
}

func (a AdvancedStudent) Practice() {
	fmt.Println(a.Name, "практикует", a.Focus, "на углубленном уровне.")
}

func main() {
	// Создаем ученика
	advancedStudent := AdvancedStudent{
		BaseStudent: BaseStudent{Name: "Мария Петрова", progress: 60},
		Skills: []Skill{
			{Name: "Speaking", Description: "Разговорный английский"},
			{Name: "Reading", Description: "Чтение текстов"},
		},
		Materials: []Material{
			{Title: "Advanced Grammar", Type: "Книга"},
			{Title: "Business English", Type: "Аудио"},
		},
		Focus: "деловой английский",
	}

	// Используем методы
	fmt.Println("Ученик:", advancedStudent.Name)
	fmt.Println("Прогресс:", advancedStudent.GetProgress(), "%")
	advancedStudent.UpdateProgress(10)
	advancedStudent.Practice()

	fmt.Println("Навыки:")
	for _, skill := range advancedStudent.Skills {
		fmt.Println("-", skill.Name, ":", skill.Description)
	}
	fmt.Println("Материалы:")
	for _, material := range advancedStudent.Materials {
		fmt.Println("-", material.Title, "(", material.Type, ")")
	}
}

/*
	Абстракция: Интерфейс Practiceable скрывает детали реализации.
	Композиция: Ученик содержит навыки и учебные материалы.

	Инкапсуляция: Прогресс ученика скрыт, но доступен через метод.
	Наследование: AdvancedStudent встраивает BaseStudent.
	Полиморфизм: Разные ученики могут практиковать навыки через общий интерфейс.
*/
