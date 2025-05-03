package main

import "fmt"

//Композиция позволяет создавать сложные объекты, комбинируя более простые.
//Например, ученик может иметь список оценок и домашних заданий.

// Оценка
type Grade struct {
	Subject string
	Value   int
}

// Домашнее задание
type Homework struct {
	Subject string
	Task    string
}

// Ученик
type Student struct {
	Name     string
	Grades   []Grade
	Homework []Homework
}

func main() {
	student := Student{
		Name: "Иван Иванов",
		Grades: []Grade{
			{Subject: "Математика", Value: 5},
			{Subject: "Физика", Value: 4},
		},
		Homework: []Homework{
			{Subject: "Математика", Task: "Решить задачу №5"},
			{Subject: "Физика", Task: "Подготовить доклад"},
		},
	}

	fmt.Println("Ученик:", student.Name)
	fmt.Println("Оценки:")
	for _, grade := range student.Grades {
		fmt.Println("-", grade.Subject, ":", grade.Value)
	}
	fmt.Println("Домашние задания:")
	for _, hw := range student.Homework {
		fmt.Println("-", hw.Subject, ":", hw.Task)
	}
}
