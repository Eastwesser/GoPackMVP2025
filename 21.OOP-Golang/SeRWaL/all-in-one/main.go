package main

import "fmt"

// Абстракция: интерфейс для пользователей
type User interface {
	Login()
	ViewGrades()
}

// Базовый класс (Инкапсуляция)
type BaseUser struct {
	Name     string
	password string
}

func (u BaseUser) GetName() string {
	return u.Name
}

func (u *BaseUser) ChangePassword(newPassword string) {
	u.password = newPassword
	fmt.Println(u.Name, "изменил пароль.")
}

// Композиция: ученик с оценками и домашними заданиями
type Grade struct {
	Subject string
	Value   int
}

type Homework struct {
	Subject string
	Task    string
}

// Наследование (встраивание) и полиморфизм
type Student struct {
	BaseUser
	Grades   []Grade
	Homework []Homework
}

func (s Student) Login() {
	fmt.Println(s.Name, "вошел в систему как ученик.")
}

func (s Student) ViewGrades() {
	fmt.Println(s.Name, "просматривает свои оценки:")
	for _, grade := range s.Grades {
		fmt.Println("-", grade.Subject, ":", grade.Value)
	}
}

func main() {
	// Создаем ученика
	student := Student{
		BaseUser: BaseUser{Name: "Иван Иванов", password: "qwerty"},
		Grades: []Grade{
			{Subject: "Математика", Value: 5},
			{Subject: "Физика", Value: 4},
		},
		Homework: []Homework{
			{Subject: "Математика", Task: "Решить задачу №5"},
			{Subject: "Физика", Task: "Подготовить доклад"},
		},
	}

	// Используем методы
	fmt.Println("Пользователь:", student.GetName())
	student.Login()
	student.ViewGrades()
	student.ChangePassword("newpassword123")
}

/*
	Абстракция: Интерфейс User скрывает детали реализации.
	Композиция: Ученик содержит оценки и домашние задания.

	Инкапсуляция: Пароль пользователя скрыт, но доступен через метод.
	Наследование: Student встраивает BaseUser. Новая структура получает базовые настройки
	Полиморфизм: Разные пользователи могут входить в систему через общий интерфейс. Разные структуры пользуются одним методом.
*/
