package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

// UserEntity is a a clean entity
type UserEntity struct {
	ID       int
	Name     string
	Age      int
	Nickname string
	Phone    string
	Email    string
}

// --- Domain Layer (Сущности) ---
// Это ядро системы, не зависит ни от чего внешнего

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required,min=2"`
	Age      int    `json:"age" validate:"gte=0,lte=130"`
	Nickname string `json:"nickname" validate:"required"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email" validate:"required,email"`
}

// --- Use Case Layer (Сценарии) ---
// Бизнес-правила приложения

// --- Infrastructure Layer (Реализация репозитория) ---
// Детали работы с данными (в данном случае in-memory, но обычно это база данных Постгрес)

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

// this func creates a user with info
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

// this func gets user's info
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

// this func updates user's info
func updateUserInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var updatedUser User

	if err = json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(updatedUser)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

// this func kills a user with entire info
func deleteUserInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "not found", http.StatusNotFound)
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

	// CRUDs
	r.Route("/user", func(r chi.Router) {
		// POST /users - создать нового пользователя
		r.Post("/", createUser)
		// GET /users/{id} - получить пользователя по ID
		r.Get("/{id}", readUserInfo)
		// PUT /users/{id} - обновить пользователя
		r.Put("/{id}", updateUserInfo)
		// DELETE /users/{id} - удалить пользователя
		r.Delete("/{id}", deleteUserInfo)

		// GET /users - список всех пользователей
		r.Get("/", listUsers)
	})

	fmt.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
