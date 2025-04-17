package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type OriginalUser struct {
	ID      int     `json:"id"`
	Email   string  `json:"email"`
	Amount  float64 `json:"amount"`
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
	Email   string  `json:"email,omitempty"`
	Amount  float64 `json:"amount"`
	Profile struct {
		Avatar    string `json:"avatar,omitempty"`
		LastName  string `json:"lastName,omitempty"`
		FirstName string `json:"firstName,omitempty"`
	} `json:"profile"`
	Username  string `json:"username,omitempty"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("http://80.000.000.00:8000/users")
	if err != nil {
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var originals []OriginalUser

	if err := json.NewDecoder(resp.Body).Decode(&originals); err != nil {
		http.Error(w, "failed decode", http.StatusInternalServerError)
		return
	}

	var modifiedUsers []ModifiedUser

	for _, original := range originals {
		var modified ModifiedUser
		modified.ID = original.ID
		modified.Amount = original.Amount
		modified.CreatedAt = original.CreatedAt
		modified.CreatedBy = original.CreatedBy

		if original.Amount <= 1000000 {
			modified.Email = original.Email
			modified.Username = original.Username
			modified.Profile.FirstName = original.Profile.FirstName
			modified.Profile.LastName = original.Profile.LastName
			modified.Profile.Avatar = original.Profile.Avatar
		}
		modifiedUsers = append(modifiedUsers, modified)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(modifiedUsers)
}

func main() {
	http.HandleFunc("/users", userHandler)
	fmt.Println("Server running on 8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("error server", err)
	}
}
