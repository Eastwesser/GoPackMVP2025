package main

import "fmt"

//В Go нет классического наследования, но можно использовать встраивание (embedding).
//Например, базовый класс "Ученик" и специализированный класс "Ученик с углубленным изучением".

// Базовый класс
type Student struct {
	Name string
}

func (s Student) Study() {
	fmt.Println(s.Name, "учит английский.")
}

// Специализированный класс
type AdvancedStudent struct {
	Student
	Focus string
}

func (a AdvancedStudent) StudyAdvanced() {
	fmt.Println(a.Name, "изучает", a.Focus, "на углубленном уровне.")
}

func main() {
	advancedStudent := AdvancedStudent{
		Student: Student{Name: "Мария Петрова"},
		Focus:   "грамматика",
	}

	advancedStudent.Study()
	advancedStudent.StudyAdvanced()
}
