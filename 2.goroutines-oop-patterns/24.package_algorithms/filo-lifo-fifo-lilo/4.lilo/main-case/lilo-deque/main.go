package main

import (
	"fmt"
)

// LILO - это дек, где заказы добавляются в начало, а извлекаются с конца

// Pharmacy represents a queue for processing orders
type Pharmacy struct {
	orders []string // Slice to store orders
}

// AddOrder adds an order to the front of the queue
func (p *Pharmacy) AddOrder(order string) {
	p.orders = append([]string{order}, p.orders...) // Add to the front
}

// ProcessNextOrder removes and returns the order from the end of the queue
func (p *Pharmacy) ProcessNextOrder() string {
	if len(p.orders) == 0 {
		return "No orders to process"
	}

	// Get the last order
	order := p.orders[len(p.orders)-1]
	// Remove the last order
	p.orders = p.orders[:len(p.orders)-1]

	return order
}

func main() {
	pharmacy := Pharmacy{}

	// Add orders to the queue
	pharmacy.AddOrder("Order 1: Paracetamol")
	pharmacy.AddOrder("Order 2: Ibuprofen")
	pharmacy.AddOrder("Order 3: Amoxicillin")

	// Process orders in LILO (Last-In, Last-Out) order
	fmt.Println(pharmacy.ProcessNextOrder()) // Output: Order 1: Paracetamol
	fmt.Println(pharmacy.ProcessNextOrder()) // Output: Order 2: Ibuprofen
	fmt.Println(pharmacy.ProcessNextOrder()) // Output: Order 3: Amoxicillin
	fmt.Println(pharmacy.ProcessNextOrder()) // Output: No orders to process
}
