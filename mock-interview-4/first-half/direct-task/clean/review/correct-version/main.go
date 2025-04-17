package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	Amount    int     `json:"amount"`
	Profile   Profile `json:"profile"`
	Password  string  `json:"password"`
	Username  string  `json:"username"`
	CreatedAt string  `json:"createdAt"`
	CreatedBy string  `json:"createdBy"`
}

type Profile struct {
	Avatar     string `json:"avatar"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	StaticData string `json:"staticData,omitempty"`
}

type RichUser struct {
	ID       int      `json:"id"`
	Email    *string  `json:"email,omitempty"`
	Amount   int      `json:"amount"`
	Username *string  `json:"username,omitempty"`
	Profile  *Profile `json:"profile,omitempty"`
}

func main() {
	http.HandleFunc("/users", duplicator)
	fmt.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func duplicator(w http.ResponseWriter, r *http.Request) {
	// Создаем кастомный клиент с настройками
	client := &http.Client{}

	// Создаем запрос
	req, err := http.NewRequest("GET", "http://80.000.000.00:8000/users", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Выполняем запрос через client.Do()
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Декодируем ответ
	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Фильтруем пользователей
	var filteredUsers []RichUser

	for _, user := range users {
		richUser := RichUser{
			ID:     user.ID,
			Amount: user.Amount,
		}

		if user.Amount <= 1000000 {
			// Копируем данные только если сумма <= 1000000
			email := user.Email
			username := user.Username
			profile := user.Profile
			profile.StaticData = "" // Очищаем staticData

			richUser.Email = &email
			richUser.Username = &username
			richUser.Profile = &profile
		}

		filteredUsers = append(filteredUsers, richUser)
	}

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(filteredUsers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
