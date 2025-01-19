package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Random integer between 1 and 100
	randInt := rand.Intn(100) + 1
	fmt.Printf("Random integer: %d\n", randInt)

	// Random float between 0.0 and 1.0
	randFloat := rand.Float64()
	fmt.Printf("Random float: %f\n", randFloat)

	// Simulate a dice roll (1 to 6)
	diceRoll := rand.Intn(6) + 1
	fmt.Printf("Dice roll: %d\n", diceRoll)

	// Randomly pick an element from a slice
	colors := []string{"Red", "Blue", "Green", "Yellow", "Purple"}
	randomColor := colors[rand.Intn(len(colors))]
	fmt.Printf("Random color: %s\n", randomColor)
}
