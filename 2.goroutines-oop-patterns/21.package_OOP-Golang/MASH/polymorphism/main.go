package main

import "fmt"

//Полиморфизм позволяет использовать разные типы через общий интерфейс.
//Например, пользователи могут входить в систему, но каждый делает это по-своему.

// Интерфейс для входа в систему
type Loginable interface {
	Login()
}

// Ученик
type Student struct {
	Name string
}

func (s Student) Login() {
	fmt.Println(s.Name, "вошел в систему как ученик.")
}

// Учитель
type Teacher struct {
	Name string
}

func (t Teacher) Login() {
	fmt.Println(t.Name, "вошел в систему как учитель.")
}

func main() {
	users := []Loginable{
		Student{Name: "Иван Иванов"},
		Teacher{Name: "Мария Петрова"},
	}

	for _, user := range users {
		user.Login()
	}
}
