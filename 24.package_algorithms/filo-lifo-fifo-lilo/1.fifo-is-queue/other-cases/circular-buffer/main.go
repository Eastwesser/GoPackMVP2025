package main

import "fmt"

// PharmacyManager определяет интерфейс для управления очередью пациентов.
type PharmacyManager interface {
	AddPatient(patient string)
	ServePatient() (string, bool)
}

// PharmacyQueue представляет очередь пациентов в аптеке.
type PharmacyQueue struct {
	queue []string
	size  int
	start int
	end   int
	count int
}

// NewPharmacyQueue создает новую очередь пациентов с заданным размером.
func NewPharmacyQueue(size int) *PharmacyQueue {
	return &PharmacyQueue{
		queue: make([]string, size),
		size:  size,
		start: 0,
		end:   0,
		count: 0,
	}
}

// AddPatient добавляет пациента в очередь.
func (pq *PharmacyQueue) AddPatient(patient string) {
	if pq.count == pq.size {
		// Очередь заполнена, удаляем самого старого пациента
		pq.start = (pq.start + 1) % pq.size
		pq.count--
	}

	pq.queue[pq.end] = patient
	pq.end = (pq.end + 1) % pq.size
	pq.count++
}

// ServePatient выдает лекарство следующему пациенту в очереди.
func (pq *PharmacyQueue) ServePatient() (string, bool) {
	if pq.count == 0 {
		return "", false // Очередь пуста
	}

	patient := pq.queue[pq.start]
	pq.start = (pq.start + 1) % pq.size
	pq.count--

	return patient, true
}

func main() {
	pharmacy := NewPharmacyQueue(3) // Очередь на 3 пациента

	pharmacy.AddPatient("Patient A")
	pharmacy.AddPatient("Patient B")
	pharmacy.AddPatient("Patient C")
	pharmacy.AddPatient("Patient D") // Заменяет Patient A

	for {
		patient, ok := pharmacy.ServePatient()
		if !ok {
			fmt.Println("No more patients in the queue.") // Выводит: No more patients in the queue.
			break
		}
		fmt.Println("Serving:", patient) // Выводит: Serving: Patient B, Serving: Patient C, Serving: Patient D
	}
}
