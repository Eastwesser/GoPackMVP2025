package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Avatar     string `json:"avatar"`
	Password   string `json:"password"`
	StaticData string `json:"staticData"`
	Balance    int    `json:"balance"`
}

type FilteredUser struct {
	ID        int     `json:"id"`
	Username  *string `json:"username,omitempty"`
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Avatar    *string `json:"avatar,omitempty"`
	Balance   int     `json:"balance"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Make request to the original API
	response, err := http.Get("https://api.example.com" + r.URL.Path)
	if err != nil {
		http.Error(w, "Failed to fetch user data", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(response.Body)

	// Check if the response is successful
	if response.StatusCode != http.StatusOK {
		http.Error(w, "User not found", response.StatusCode)
		return
	}

	// Decode the user data
	var user User
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse user data", http.StatusInternalServerError)
		return
	}

	// Create filtered response
	filtered := FilteredUser{
		ID:      user.ID,
		Balance: user.Balance,
	}

	// Only include sensitive fields if balance is <= 50000
	if user.Balance <= 50000 {
		filtered.Username = &user.Username
		filtered.Email = &user.Email
		filtered.FirstName = &user.FirstName
		filtered.LastName = &user.LastName
		filtered.Avatar = &user.Avatar
	}

	// Set content type and encode the response
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(filtered)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/users/", userHandler)

	// Start the server
	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
