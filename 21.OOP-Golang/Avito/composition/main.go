package main

import "fmt"

//Композиция позволяет создавать сложные объекты, комбинируя более простые.
//Например, пользователь Avito может иметь профиль и список объявлений.

// Профиль пользователя
type Profile struct {
	Name  string
	Email string
}

// Объявление
type Advertisement struct {
	Title string
	Price float64
}

// Пользователь
type User struct {
	Profile        Profile
	Advertisements []Advertisement
}

func main() {
	user := User{
		Profile: Profile{Name: "Иван Иванов", Email: "ivan@example.com"},
		Advertisements: []Advertisement{
			{Title: "Продам велосипед", Price: 15000},
			{Title: "Сдам гараж", Price: 5000},
		},
	}

	fmt.Println("Пользователь:", user.Profile.Name)
	fmt.Println("Email:", user.Profile.Email)
	fmt.Println("Объявления:")
	for _, ad := range user.Advertisements {
		fmt.Println("-", ad.Title, "Цена:", ad.Price, "руб.")
	}
}
