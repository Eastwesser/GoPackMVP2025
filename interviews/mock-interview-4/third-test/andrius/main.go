package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type User struct {
	ID       int
	Name     string
	Age      int
	Nickname string
	Phone    string
	Email    string
}

var users []User
var atmID = 1

type IUser interface {
	ID() int
}

func NewUser(id int, name string, age int, nickname string, email string) []User {

	user := User{
		ID:       id,
		Name:     name,
		Age:      age,
		Nickname: nickname,
		Email:    email,
	}
	users = append(users, user)

	return users
}

func createUser() {
	panic("implement me")
}

func readUserInfo() {
	panic("implement me")
}

func updateUserInfo() {
	panic("implement me")
}

func deleteUserInfo() {
	panic("implement me")
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	json.

}

func main() {
	r := chi.NewRouter()

	// init test data
	users = append(users, User{
		ID:       1,
		Name:     "John Doe",
		Age:      20,
		Nickname: "John Doe",
		Email:    "john@doe.com",
		Phone:    "0987654321",
	})
	atmID++

	r.Route("/user", func(r chi.Router) {
		// CRUDs
		r.Post("/", createUser)
		r.Get("/{id}", readUserInfo)
		r.Put("/{id}", updateUserInfo)
		r.Delete("/{id}", deleteUserInfo)
		// get all users
		r.Get("/", listUsers)
	})

	fmt.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
