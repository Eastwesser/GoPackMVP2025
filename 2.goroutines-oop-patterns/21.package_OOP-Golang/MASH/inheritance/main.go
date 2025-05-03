package main

import "fmt"

// В Go нет классического наследования, но можно использовать встраивание (embedding).
//Например, базовый класс "Пользователь" и специализированный класс "Администратор".

// Базовый класс
type User struct {
	Name string
}

func (u User) Login() {
	fmt.Println(u.Name, "вошел в систему.")
}

// Специализированный класс
type Admin struct {
	User
	Role string
}

func (a Admin) ManageSystem() {
	fmt.Println(a.Name, "управляет системой как", a.Role)
}

func main() {
	admin := Admin{
		User: User{Name: "Алексей Сидоров"},
		Role: "Администратор",
	}

	admin.Login()
	admin.ManageSystem()
}
