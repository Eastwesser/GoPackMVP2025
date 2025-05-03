package main

import "fmt"

//Полиморфизм позволяет использовать разные типы через общий интерфейс.
//Например, ученики могут практиковать навыки, но каждый делает это по-своему.

// Интерфейс для практики
type Practiceable interface {
	Practice()
}

// Ученик начальных классов
type ElementaryStudent struct {
	Name string
}

func (e ElementaryStudent) Practice() {
	fmt.Println(e.Name, "практикует базовые навыки.")
}

// Ученик старших классов
type HighSchoolStudent struct {
	Name string
}

func (h HighSchoolStudent) Practice() {
	fmt.Println(h.Name, "практикует продвинутые навыки.")
}

func main() {
	students := []Practiceable{
		ElementaryStudent{Name: "Иван Иванов"},
		HighSchoolStudent{Name: "Мария Петрова"},
	}

	for _, student := range students {
		student.Practice()
	}
}
