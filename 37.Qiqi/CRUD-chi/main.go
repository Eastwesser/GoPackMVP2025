package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users      = make(map[int]User)
	usersMutex sync.RWMutex
	idCounter  = 1
)

func main() {
	r := chi.NewRouter()

	// Маршруты для пользователей
	r.Route("/users", func(r chi.Router) {
		r.Get("/", listUsers)   // GET /users
		r.Post("/", createUser) // POST /users

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", getUser)       // GET /users/{id}
			r.Put("/", updateUser)    // PUT /users/{id}
			r.Delete("/", deleteUser) // DELETE /users/{id}
		})
	})

	// Запуск сервера
	http.ListenAndServe(":8080", r)
}

// listUsers возвращает список всех пользователей
func listUsers(w http.ResponseWriter, r *http.Request) {
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	userList := make([]User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}

	respondWithJSON(w, http.StatusOK, userList)
}

// createUser создает нового пользователя
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	newUser.ID = idCounter
	users[newUser.ID] = newUser
	idCounter++

	respondWithJSON(w, http.StatusCreated, newUser)
}

// getUser возвращает пользователя по ID
func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	usersMutex.RLock()
	defer usersMutex.RUnlock()

	user, exists := users[id]
	if !exists {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// updateUser обновляет данные пользователя
func updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	if _, exists := users[id]; !exists {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	updatedUser.ID = id
	users[id] = updatedUser

	respondWithJSON(w, http.StatusOK, updatedUser)
}

// deleteUser удаляет пользователя
func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	if _, exists := users[id]; !exists {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	delete(users, id)
	w.WriteHeader(http.StatusNoContent)
}

// Вспомогательные функции для ответов
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
