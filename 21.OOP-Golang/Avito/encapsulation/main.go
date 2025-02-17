package main

import "fmt"

//Инкапсуляция позволяет скрыть внутренние детали объекта.
//Например, баланс пользователя может быть скрыт, но доступен через метод.

type User struct {
	name    string
	balance float64
}

func (u User) GetBalance() float64 {
	return u.balance
}

func (u *User) AddFunds(amount float64) {
	if amount > 0 {
		u.balance += amount
		fmt.Println(u.name, "пополнил баланс на", amount, "руб.")
	}
}

func main() {
	user := User{name: "Иван Иванов", balance: 1000}
	fmt.Println("Баланс", user.name, ":", user.GetBalance(), "руб.")

	user.AddFunds(500)
	fmt.Println("Новый баланс:", user.GetBalance(), "руб.")
}
