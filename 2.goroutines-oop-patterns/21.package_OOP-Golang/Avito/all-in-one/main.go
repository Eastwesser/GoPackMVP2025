package main

import "fmt"

// Абстракция: интерфейс для публикации
type Publisher interface {
	Publish()
}

// Базовый класс (Инкапсуляция)
type User struct {
	Name    string
	balance float64
}

func (u User) GetBalance() float64 {
	return u.balance
}

func (u *User) AddFunds(amount float64) {
	if amount > 0 {
		u.balance += amount
		fmt.Println(u.Name, "пополнил баланс на", amount, "руб.")
	}
}

// Композиция: пользователь с объявлениями
type Advertisement struct {
	Title string
	Price float64
}

// Наследование (встраивание) и полиморфизм
type Seller struct {
	User
	Advertisements []Advertisement
	Rating         float64
}

func (s Seller) Publish() {
	fmt.Println(s.Name, "опубликовал новое объявление.")
}

func main() {
	// Создаем продавца
	seller := Seller{
		User: User{Name: "Иван Иванов", balance: 1000},
		Advertisements: []Advertisement{
			{Title: "Продам ноутбук", Price: 50000},
			{Title: "Продам телефон", Price: 20000},
		},
		Rating: 4.8,
	}

	// Используем методы
	fmt.Println("Баланс", seller.Name, ":", seller.GetBalance(), "руб.")
	seller.AddFunds(500)
	seller.Publish()

	fmt.Println("Объявления:")
	for _, ad := range seller.Advertisements {
		fmt.Println("-", ad.Title, "Цена:", ad.Price, "руб.")
	}
}

/*
	Абстракция: Интерфейс Publisher скрывает детали реализации.
	Композиция: Пользователь содержит профиль и список объявлений.

	Инкапсуляция: Баланс пользователя скрыт, но доступен через метод.
	Наследование: Seller встраивает User.
	Полиморфизм: Разные пользователи могут публиковать объявления через общий интерфейс.
*/
