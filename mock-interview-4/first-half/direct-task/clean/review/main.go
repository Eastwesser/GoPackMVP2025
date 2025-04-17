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

// RICHNESS FILTER
type RichUser struct {
	ID       int      `json:"id"`
	Email    *string  `json:"email,omitempty"`
	Amount   int      `json:"avatar"`
	Username *string  `json:"username,omitempty"`
	Profile  *Profile `json:"profile,omitempty"`
}

func main() {
	http.HandleFunc("/users", duplicator)
	fmt.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func duplicator(w http.ResponseWriter, r *http.Request) {

	clientOfThisTaskXDDD := &http.Client{}

	req, err := http.NewRequest(r.Method, "http://80.000.000.00:8000/users", nil)
	if err != nil {
		fmt.Println(err)
	}

	response, err := clientOfThisTaskXDDD.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	//response, _ := http.Get("http://80.000.000.00:8000/users")
	//defer response.Body.Close()
	////
	//var users []User
	//json.NewDecoder(response.Body).Decode(&users)

	var users []User
	json.NewDecoder(response.Body).Decode(&users)

	// user filtering
	var userFilterList []RichUser

	for _, user := range users {
		//filteredUsers := RichUser{
		//	ID:     user.ID,
		//	Amount: user.Amount,
		//}
		//
		//if user.Amount <= 1000000 {
		//	filteredUsers.Email = &user.Email
		//	filteredUsers.Username = &user.Username
		//	richProfile := user.Profile
		//	richProfile.StaticData = ""
		//	filteredUsers.Profile = &richProfile
		//}

		filteredUsers := RichUser{
			ID:     user.ID,
			Amount: user.Amount,
		}

		if user.Amount <= 1000000 {
			email := user.Email
			username := user.Username
			profile := user.Profile
			profile.StaticData = ""

			filteredUsers.Email = &email
			filteredUsers.Username = &username
			filteredUsers.Profile = &profile
		}

		userFilterList = append(userFilterList, filteredUsers)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userFilterList)
}
