package main

import (
	"fmt"
)

func main() {
	// Map to store pills and their purposes
	pills := map[string]string{
		"Headache":     "Nurofen",
		"Stomachache":  "Omeprazole",
		"Skin Issues":  "Hydrocortisone Cream",
		"Anti-Allergy": "Cetirizine",
		"Fever":        "Paracetamol",
		"Muscle Pain":  "Ibuprofen",
		"Indigestion":  "Antacid Tablets",
		"Cold & Flu":   "Lemsip",
	}

	// Display available options
	fmt.Println("Welcome to your personal pill reminder!")
	fmt.Println("Choose a category to see the recommended pill:")
	fmt.Println("1. Headache")
	fmt.Println("2. Stomachache")
	fmt.Println("3. Skin Issues")
	fmt.Println("4. Anti-Allergy")
	fmt.Println("5. Fever")
	fmt.Println("6. Muscle Pain")
	fmt.Println("7. Indigestion")
	fmt.Println("8. Cold & Flu")

	// Get user input
	var choice int
	fmt.Print("Enter the number of your choice: ")
	fmt.Scan(&choice)

	// Use switch-case to handle user input
	switch choice {
	case 1:
		fmt.Printf("For Headache: %s\n", pills["Headache"])
	case 2:
		fmt.Printf("For Stomachache: %s\n", pills["Stomachache"])
	case 3:
		fmt.Printf("For Skin Issues: %s\n", pills["Skin Issues"])
	case 4:
		fmt.Printf("For Anti-Allergy: %s\n", pills["Anti-Allergy"])
	case 5:
		fmt.Printf("For Fever: %s\n", pills["Fever"])
	case 6:
		fmt.Printf("For Muscle Pain: %s\n", pills["Muscle Pain"])
	case 7:
		fmt.Printf("For Indigestion: %s\n", pills["Indigestion"])
	case 8:
		fmt.Printf("For Cold & Flu: %s\n", pills["Cold & Flu"])
	default:
		fmt.Println("Invalid choice! Please select a number between 1 and 8.")
	}
}
