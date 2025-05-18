package main

import "github.com/go-chi/chi/v5"

type User struct {
	ID       int
	Name     string
	Age      int
	Nickname string
	Phone    string
	Email    string
}

type Users []User

var users Users

type IUser interface {
	ID() int
}

func NewUser(id int, name string, age int, nickname string, email string) IUser {

	user := User{
		ID:       id,
		Name:     name,
		Age:      age,
		Nickname: nickname,
		Email:    email,
	}
	users = append(users, user)

	return user
}

func CreateUser(u *IUser) {
	panic("implement me")
}

func ReadUserInfo() {
	panic("implement me")
}

func UpdateUserInfo() {
	panic("implement me")
}

func DeleteUserInfo() {
	panic("implement me")
}

func main() {
	CreateUser()
	ReadUserInfo()
	UpdateUserInfo()
	DeleteUserInfo()

	r := chi.NewRouter()
	r
}
