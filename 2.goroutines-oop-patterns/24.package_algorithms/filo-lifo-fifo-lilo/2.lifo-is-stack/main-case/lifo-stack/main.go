package main

import (
	"fmt"
)

// LIFO - это стек, где последнее добавленное лекарство будет первым извлеченным

type Warehouse struct {
	stack []string
}

func (w *Warehouse) AddMedicine(medicine string) {
	w.stack = append(w.stack, medicine)
}

func (w *Warehouse) RetrieveMedicine() string {
	if len(w.stack) == 0 {
		return "No medicines in the warehouse"
	}

	index := len(w.stack) - 1
	medicine := w.stack[index]
	w.stack = w.stack[:index]

	return medicine
}

func main() {
	warehouse := Warehouse{}

	warehouse.AddMedicine("Medicine 1: Paracetamol")
	warehouse.AddMedicine("Medicine 2: Ibuprofen")
	warehouse.AddMedicine("Medicine 3: Amoxicillin")

	fmt.Println(warehouse.RetrieveMedicine()) // Output: Medicine 3: Amoxicillin
	fmt.Println(warehouse.RetrieveMedicine()) // Output: Medicine 2: Ibuprofen
	fmt.Println(warehouse.RetrieveMedicine()) // Output: Medicine 1: Paracetamol
	fmt.Println(warehouse.RetrieveMedicine()) // Output: No medicines in the warehouse
}
