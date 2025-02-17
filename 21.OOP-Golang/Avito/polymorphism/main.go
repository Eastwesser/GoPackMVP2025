package main

import "fmt"

//Полиморфизм позволяет использовать разные типы через общий интерфейс.
//Например, пользователи могут публиковать объявления, но каждый делает это по-своему.

// Интерфейс для публикации
type Publisher interface {
	Publish()
}

// Пользователь
type User struct {
	Name string
}

func (u User) Publish() {
	fmt.Println(u.Name, "опубликовал объявление.")
}

// Администратор
type Admin struct {
	Name string
}

func (a Admin) Publish() {
	fmt.Println(a.Name, "опубликовал служебное объявление.")
}

func main() {
	publishers := []Publisher{
		User{Name: "Иван Иванов"},
		Admin{Name: "Администратор"},
	}

	for _, publisher := range publishers {
		publisher.Publish()
	}
}
