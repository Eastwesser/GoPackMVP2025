package main

import (
	"fmt"
)

// FILO - это стек, где последний добавленный рецепт будет первым обработанным

type Pharmacy struct {
	stack []string
}

func (p *Pharmacy) AddPrescription(prescription string) {
	p.stack = append(p.stack, prescription)
}

func (p *Pharmacy) ProcessNextPrescription() string {
	if len(p.stack) == 0 {
		return "No prescriptions to process"
	}

	index := len(p.stack) - 1
	prescription := p.stack[index]
	p.stack = p.stack[:index]

	return prescription
}

func main() {
	pharmacy := Pharmacy{}

	pharmacy.AddPrescription("Prescription 1: Paracetamol")
	pharmacy.AddPrescription("Prescription 2: Ibuprofen")
	pharmacy.AddPrescription("Prescription 3: Amoxicillin")

	fmt.Println(pharmacy.ProcessNextPrescription()) // Output: Prescription 3: Amoxicillin
	fmt.Println(pharmacy.ProcessNextPrescription()) // Output: Prescription 2: Ibuprofen
	fmt.Println(pharmacy.ProcessNextPrescription()) // Output: Prescription 1: Paracetamol
	fmt.Println(pharmacy.ProcessNextPrescription()) // Output: No prescriptions to process
}
