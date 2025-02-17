package main

import "fmt"

//Инкапсуляция позволяет скрыть внутренние детали объекта.
//Например, пароль пользователя может быть скрыт, но доступен через метод.

type User struct {
	name     string
	password string
}

func (u User) GetName() string {
	return u.name
}

func (u *User) ChangePassword(newPassword string) {
	u.password = newPassword
	fmt.Println(u.name, "изменил пароль.")
}

func main() {
	user := User{name: "Иван Иванов", password: "qwerty"}
	fmt.Println("Пользователь:", user.GetName())

	user.ChangePassword("newpassword123")
}
