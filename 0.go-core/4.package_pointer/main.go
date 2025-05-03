package main

import "fmt"

func increment(ptr *int) {
	*ptr++
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 42
	fmt.Println("Original value of x:", x)

	increment(&x)
	fmt.Println("Value of x after increment:", x)

	// Passing pointers to swap function
	a, b := 10, 20
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a = %d, b = %d\n", a, b)
}
