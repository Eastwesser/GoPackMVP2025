package main

import (
	"fmt"
	"github.com/golang-collections/collections/deque"
)

// LILO - это дек, где заказы добавляются в начало, а извлекаются с конца

type Pharmacy struct {
	deque *deque.Deque
}

func (p *Pharmacy) AddOrder(order string) {
	p.deque.PushFront(order)
}

func (p *Pharmacy) ProcessNextOrder() string {
	if p.deque.Len() == 0 {
		return "No orders to process"
	}

	order := p.deque.PopBack().(string)
	return order
}

func main() {
	pharmacy := Pharmacy{deque: deque.New()}

	pharmacy.AddOrder("Order 1: Paracetamol")
	pharmacy.AddOrder("Order 2: Ibuprofen")
	pharmacy.AddOrder("Order 3: Amoxicillin")

	fmt.Println(pharmacy.ProcessNextOrder()) // Output: Order 1: Paracetamol
	fmt.Println(pharmacy.ProcessNextOrder()) // Output: Order 2: Ibuprofen
	fmt.Println(pharmacy.ProcessNextOrder()) // Output: Order 3: Amoxicillin
	fmt.Println(pharmacy.ProcessNextOrder()) // Output: No orders to process
}
