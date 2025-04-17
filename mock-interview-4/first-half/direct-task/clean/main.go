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
	StaticData string `json:"staticData,omitempty"` // StaticData явно очищается (richProfile.StaticData = "") и не попадает в JSON благодаря omitempty в Profile.
}

// ЭТО СТРУКТУРА - ФИЛЬТР ПО БОГАТСТВУ. Пока * не ставим, тестим, потом ставим
type RichUser struct {
	ID       int      `json:"id"`
	Email    *string  `json:"email,omitempty"`
	Amount   int      `json:"amount"`
	Username *string  `json:"username,omitempty"`
	Profile  *Profile `json:"profile,omitempty"`
}

//nil → поля нет в JSON (благодаря omitempty).

//Указатели нужны только для полей, которые могут скрываться.
//Для ID и Amount указатели избыточны (и даже вредны).
//Без указателей omitempty не работает как ожидается.

//*string → поле есть, даже если это пустая строка.

func main() {
	http.HandleFunc("/users", mimic)
	fmt.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func mimic(w http.ResponseWriter, r *http.Request) {
	// Используем http.Get для получения данных с исходного API.
	response, _ := http.Get("http://80.000.000.00:8000/users")
	defer response.Body.Close()

	var users []User
	json.NewDecoder(response.Body).Decode(&users)

	// FILTER
	var userFilterList []RichUser

	for _, user := range users {
		filteredUsers := RichUser{
			ID:     user.ID,
			Amount: user.Amount,
		}
		if user.Amount <= 50000 {
			filteredUsers.Email = &user.Email
			filteredUsers.Username = &user.Username
			richProfile := user.Profile
			// Нам не нужна **staticData** Create filtered profile without staticData!!!
			richProfile.StaticData = "" // пустое значение, но оно не пустое для omitempty, nil поля нет в JSON
			filteredUsers.Profile = &richProfile
		}

		userFilterList = append(userFilterList, filteredUsers)
	}

	// Set content type and encode the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userFilterList)
}
