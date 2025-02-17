package main

import "fmt"

type Student struct {
	name     string
	progress int
}

func (s Student) GetProgress() int {
	return s.progress
}

func (s *Student) UpdateProgress(amount int) {
	if amount > 0 {
		s.progress += amount
		fmt.Println(s.name, "улучшил прогресс на", amount, "%")
	}
}

func main() {
	student := Student{name: "Иван Иванов", progress: 50}
	fmt.Println("Прогресс", student.name, ":", student.GetProgress(), "%")

	student.UpdateProgress(10)
	fmt.Println("Новый прогресс:", student.GetProgress(), "%")
}
