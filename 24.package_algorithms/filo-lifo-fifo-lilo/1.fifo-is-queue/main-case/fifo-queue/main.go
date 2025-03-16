package main

import (
	"fmt"
)

// FIFO - это очередь, в которой рецепты добавляются в конец, а извлекаются из начала

type Pharmacy struct {
	queue []string
}

func (p *Pharmacy) AddPrescription(prescription string) {
	p.queue = append(p.queue, prescription)
}

func (p *Pharmacy) ProcessNextPrescription() string {
	if len(p.queue) == 0 {
		return "No prescriptions to process"
	}

	prescription := p.queue[0]
	p.queue = p.queue[1:]

	return prescription
}

func main() {
	pharmacy := Pharmacy{}

	pharmacy.AddPrescription("Prescription 1: Paracetamol")
	pharmacy.AddPrescription("Prescription 2: Ibuprofen")
	pharmacy.AddPrescription("Prescription 3: Amoxicillin")

	fmt.Println(pharmacy.ProcessNextPrescription()) // Output: Prescription 1: Paracetamol
	fmt.Println(pharmacy.ProcessNextPrescription()) // Output: Prescription 2: Ibuprofen
	fmt.Println(pharmacy.ProcessNextPrescription()) // Output: Prescription 3: Amoxicillin
	fmt.Println(pharmacy.ProcessNextPrescription()) // Output: No prescriptions to process
}
