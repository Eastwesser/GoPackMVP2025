package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Структура сотрудника
type Employee struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Phone string `json:"phone,omitempty"` // omitempty - телефон не обязателен
}

// Хранилище сотрудников
type EmployeeStorage struct {
	mu        sync.Mutex          // Для потокобезопасности
	employees map[string]Employee // Мапа для хранения
}

// Добавление сотрудника
func (s *EmployeeStorage) Add(emp Employee) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.employees[emp.ID]; exists {
		return fmt.Errorf("employee with ID %s already exists", emp.ID)
	}

	// Простые проверки
	if emp.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if emp.Year < 1900 || emp.Year > 2100 {
		return fmt.Errorf("invalid year")
	}

	s.employees[emp.ID] = emp
	return nil
}

// Получение всех сотрудников
func (s *EmployeeStorage) GetAll() []Employee {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]Employee, 0, len(s.employees))
	for _, emp := range s.employees {
		result = append(result, emp)
	}

	return result
}

// Обработчик добавления
func addEmployeeHandler(storage *EmployeeStorage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if err := storage.Add(emp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Employee %s added", emp.ID)
	}

}

// Обработчик получения всех
func getAllEmployeesHandler(storage *EmployeeStorage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		employees := storage.GetAll()
		err := json.NewEncoder(w).Encode(employees)
		if err != nil {
			return
		}
	}

}

func RunServer() {

	storage := &EmployeeStorage{
		employees: make(map[string]Employee),
	}

	// Настройка маршрутов
	http.HandleFunc("/employees/add", addEmployeeHandler(storage))
	http.HandleFunc("/employees/all", getAllEmployeesHandler(storage))

	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

func main() {
	RunServer()
}
