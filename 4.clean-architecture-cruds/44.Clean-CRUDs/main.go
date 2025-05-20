package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
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

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	user.ID = atmID
	atmID++
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
	}
}

func readUserInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(user)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	user := User{
		ID: id,
	}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
	http.Error(w, "not found", http.StatusNotFound)
}

func updateUserInfo(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func deleteUserInfo(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// this func lists all users
func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
	}
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
