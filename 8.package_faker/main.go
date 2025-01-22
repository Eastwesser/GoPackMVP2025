package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	cities := []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix"}

	// Generate a fake user profile with math rand
	name := names[rand.Intn(len(names))]
	age := rand.Intn(50) + 18
	city := cities[rand.Intn(len(cities))]

	fmt.Printf("Fake Profile: Name = %s, Age = %d, City = %s\n", name, age, city)
}
