package main

import "fmt"

//Абстракция позволяет скрыть сложность и показать только необходимые детали.
//В МЭШ это можно представить как общий интерфейс для всех пользователей, которые могут взаимодействовать с системой.

// Абстракция: интерфейс User
type User interface {
	Login()
	ViewGrades()
}

// Структура для ученика
type Student struct {
	Name string
}

func (s Student) Login() {
	fmt.Println(s.Name, "вошел в систему как ученик.")
}

func (s Student) ViewGrades() {
	fmt.Println(s.Name, "просматривает свои оценки.")
}

// Структура для учителя
type Teacher struct {
	Name string
}

func (t Teacher) Login() {
	fmt.Println(t.Name, "вошел в систему как учитель.")
}

func (t Teacher) ViewGrades() {
	fmt.Println(t.Name, "просматривает оценки учеников.")
}

func main() {
	var user User

	user = Student{Name: "Иван Иванов"}
	user.Login()
	user.ViewGrades()

	user = Teacher{Name: "Мария Петрова"}
	user.Login()
	user.ViewGrades()
}
