package main

import "fmt"

//В Go нет классического наследования, но можно использовать встраивание (embedding).
//Например, базовый класс "Пользователь" и специализированный класс "Продавец".

// Базовый класс
type User struct {
	Name string
}

func (u User) Login() {
	fmt.Println(u.Name, "вошел в систему.")
}

// Специализированный класс
type Seller struct {
	User
	Rating float64
}

func (s Seller) Sell() {
	fmt.Println(s.Name, "продает товар с рейтингом", s.Rating)
}

func main() {
	seller := Seller{
		User:   User{Name: "Иван Иванов"},
		Rating: 4.8,
	}

	seller.Login()
	seller.Sell()
}
