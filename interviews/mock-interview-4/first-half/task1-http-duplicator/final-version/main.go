package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type OriginalUser struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Amount  int    `json:"amount"`
	Profile struct {
		Avatar     string `json:"avatar"`
		LastName   string `json:"lastName"`
		FirstName  string `json:"firstName"`
		StaticData string `json:"staticData"`
	} `json:"profile"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
}

type ModifiedUser struct {
	ID      int     `json:"id"`
	Email   *string `json:"email,omitempty"`
	Amount  int     `json:"amount"`
	Profile *struct {
		Avatar    string `json:"avatar,omitempty"`
		LastName  string `json:"lastName,omitempty"`
		FirstName string `json:"firstName,omitempty"`
	} `json:"profile,omitempty"`
	Username  *string `json:"username,omitempty"`
	CreatedAt string  `json:"createdAt"`
	CreatedBy string  `json:"createdBy"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("http://83.136.232.77:8091/users")
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var originals []OriginalUser
	if err := json.NewDecoder(resp.Body).Decode(&originals); err != nil {
		http.Error(w, "Failed to decode users", http.StatusInternalServerError)
		return
	}

	var modifiedUsers []ModifiedUser

	for _, original := range originals {
		modified := ModifiedUser{
			ID:        original.ID,
			Amount:    original.Amount,
			CreatedAt: original.CreatedAt,
			CreatedBy: original.CreatedBy,
		}

		if original.Amount <= 50000 {
			email := original.Email
			username := original.Username
			modified.Email = &email
			modified.Username = &username

			modified.Profile = &struct {
				Avatar    string `json:"avatar,omitempty"`
				LastName  string `json:"lastName,omitempty"`
				FirstName string `json:"firstName,omitempty"`
			}{
				Avatar:    original.Profile.Avatar,
				LastName:  original.Profile.LastName,
				FirstName: original.Profile.FirstName,
			}
		}

		modifiedUsers = append(modifiedUsers, modified)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(modifiedUsers); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/users", userHandler)
	fmt.Println("Server running on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
