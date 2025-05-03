package main

import (
	"fmt"
)

// User struct
type User struct {
	ID   int
	Name string
}

// Простое хранилище в памяти
var (
	users  = make(map[int]User)
	nextID = 1
)

func main() {

	// C
	id := createUser("Alice")

	// R
	user, _ := readUser(id)
	fmt.Println("Created:", user)

	// U
	updateUser(id, "Alice Updated")
	user, _ = readUser(id)
	fmt.Println("Updated:", user)

	// D
	deleteUser(id)
	_, found := readUser(id)
	if !found {
		fmt.Println("User deleted successfully")
	}

}

// Create
func createUser(name string) int {

	user := User{
		ID:   nextID,
		Name: name,
	}

	users[nextID] = user
	nextID++

	return user.ID
}

// Read
func readUser(id int) (User, bool) {

	user, found := users[id]

	return user, found
}

// Update
func updateUser(id int, name string) bool {

	if user, ok := users[id]; ok {
		user.Name = name
		users[id] = user
		return true
	}

	return false
}

// Delete
func deleteUser(id int) bool {

	if _, ok := users[id]; ok {
		delete(users, id)
		return true
	}

	return false
}
