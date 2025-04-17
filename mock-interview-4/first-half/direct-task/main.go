package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Profile struct {
	Avatar     string `json:"avatar"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	StaticData string `json:"staticData,omitempty"` // будем игнорировать это поле
}

type User struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	Amount    int     `json:"amount"`
	Profile   Profile `json:"profile"`
	Password  string  `json:"password,omitempty"` // будем игнорировать это поле
	Username  string  `json:"username"`
	CreatedAt string  `json:"createdAt"`
	CreatedBy string  `json:"createdBy"`
}

type FilteredUser struct {
	ID       int      `json:"id"`
	Email    *string  `json:"email,omitempty"`
	Amount   int      `json:"amount"`
	Username *string  `json:"username,omitempty"`
	Profile  *Profile `json:"profile,omitempty"`
}

func main() {
	http.HandleFunc("/users", usersHandler) // usersHandler можно назвать как угодно
	// Start the server
	fmt.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", nil)
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	fmt.Println("Error starting server:", err)
	//}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// Make request to the original API
	//resp, err := http.Get("http://80.000.000.00:8000/users")
	//if err != nil {
	//	http.Error(w, "Failed to fetch users data", http.StatusInternalServerError)
	//	return
	//}
	//defer resp.Body.Close()

	resp, _ := http.Get("http://80.000.000.00:8000/users")
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get users", resp.StatusCode)
		return
	}

	// Decode the users data
	var users []User

	json.NewDecoder(resp.Body).Decode(&users)
	//err = json.NewDecoder(resp.Body).Decode(&users)
	//if err != nil {
	//	http.Error(w, "Failed to parse users data", http.StatusInternalServerError)
	//	return
	//}

	// Process and filter users
	var filteredUsers []FilteredUser

	for _, user := range users {
		filtered := FilteredUser{
			ID:     user.ID,
			Amount: user.Amount,
		}

		// Only include sensitive fields if amount is <= 1000000
		if user.Amount <= 1000000 {
			filtered.Email = &user.Email
			filtered.Username = &user.Username

			// Create filtered profile without staticData
			filteredProfile := Profile{
				Avatar:    user.Profile.Avatar,
				LastName:  user.Profile.LastName,
				FirstName: user.Profile.FirstName,
			}
			filtered.Profile = &filteredProfile
		}

		filteredUsers = append(filteredUsers, filtered)
	}

	// Set content type and encode the response
	w.Header().Set("Content-Type", "application/json")

	//err = json.NewEncoder(w).Encode(filteredUsers)
	//if err != nil {
	//	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	//	return
	//}
	json.NewEncoder(w).Encode(filteredUsers)
}
