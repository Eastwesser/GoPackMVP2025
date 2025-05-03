package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type OriginalUser struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	Amount    float64 `json:"amount"`
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
	StaticData string `json:"staticData"`
}

type ModifiedUser struct {
	ID        int       `json:"id"`
	Email     string    `json:"email,omitempty"`
	Amount    float64   `json:"amount"`
	Profile   *MProfile `json:"profile,omitempty"`
	Username  string    `json:"username,omitempty"`
	CreatedAt string    `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
}

type MProfile struct {
	Avatar    string `json:"avatar,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	FirstName string `json:"firstName,omitempty"`
}

func handleUser(w http.ResponseWriter, r *http.Request) {

	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", "http://80.000.000.00:8000/users", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "API request failed", resp.StatusCode)
		return
	}

	var origins []OriginalUser
	if err := json.NewDecoder(resp.Body).Decode(&origins); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	modifiedUsers := make([]ModifiedUser, 0, len(origins))
	for _, origin := range origins {
		user := ModifiedUser{
			ID:        origin.ID,
			Amount:    origin.Amount,
			CreatedAt: origin.CreatedAt,
			CreatedBy: origin.CreatedBy,
			Profile:   nil,
		}

		if origin.Amount <= 50000 {
			user.Email = origin.Email
			user.Username = origin.Username
			user.Profile = &MProfile{
				Avatar:    origin.Profile.Avatar,
				LastName:  origin.Profile.LastName,
				FirstName: origin.Profile.FirstName,
			}
		}

		modifiedUsers = append(modifiedUsers, user)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(modifiedUsers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/users", handleUser)
	fmt.Println("Server running on :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
