package main // Основной пакет приложения

import (
	"bytes"         // Для работы с буферами байтов
	"context"       // Для управления контекстами запросов
	"encoding/json" // Для работы с JSON
	"errors"        // Для создания ошибок
	"fmt"           // Для форматированного вывода
	"log"           // Для логирования
	"net/http"      // Для HTTP-сервера и клиента
	"sync"          // Для синхронизации горутин
	"time"          // Для работы с временем
)

// Employee - структура данных сотрудника (Domain Layer)
type Employee struct {
	ID    string `json:"id"`              // Уникальный идентификатор
	Name  string `json:"name"`            // Имя сотрудника
	Year  int    `json:"year"`            // Год рождения
	Phone string `json:"phone,omitempty"` // Телефон (опционально)
}

// EmployeeRepository - интерфейс хранилища (Repository Port)
type EmployeeRepository interface {
	Add(ctx context.Context, employee Employee) error // Добавить сотрудника
	GetAll(ctx context.Context) ([]Employee, error)   // Получить всех
}

// InMemoryEmployeeRepository - реализация хранилища в памяти
type InMemoryEmployeeRepository struct {
	employees map[string]Employee // Мапа для хранения
	mu        sync.RWMutex        // RWMutex для потокобезопасности
}

// NewInMemoryEmployeeRepository - конструктор хранилища
func NewInMemoryEmployeeRepository() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{
		employees: make(map[string]Employee), // Инициализация мапы
	}
}

// Add - добавление сотрудника с проверкой уникальности ID
func (r *InMemoryEmployeeRepository) Add(ctx context.Context, employee Employee) error {
	r.mu.Lock()         // Блокировка на запись
	defer r.mu.Unlock() // Гарантированное освобождение

	select {
	case <-ctx.Done(): // Проверка отмены контекста
		return ctx.Err()
	default:
		if _, exists := r.employees[employee.ID]; exists {
			return errors.New("employee already exists")
		}
		r.employees[employee.ID] = employee
		return nil
	}
}

// GetAll - получение всех сотрудников
func (r *InMemoryEmployeeRepository) GetAll(ctx context.Context) ([]Employee, error) {
	r.mu.RLock()         // Блокировка на чтение
	defer r.mu.RUnlock() // Гарантированное освобождение

	select {
	case <-ctx.Done(): // Проверка отмены контекста
		return nil, ctx.Err()
	default:
		result := make([]Employee, 0, len(r.employees))
		for _, emp := range r.employees {
			result = append(result, emp)
		}
		return result, nil
	}
}

// EmployeeUseCase - интерфейс бизнес-логики (Use Case Port)
type EmployeeUseCase interface {
	AddEmployee(ctx context.Context, id, name, phone string, year int) error
	GetAllEmployees(ctx context.Context) ([]Employee, error)
}

// EmployeeService - реализация бизнес-логики
type EmployeeService struct {
	repo EmployeeRepository // Зависимость от репозитория
}

// NewEmployeeService - конструктор сервиса
func NewEmployeeService(repo EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

// AddEmployee - добавление с валидацией
func (s *EmployeeService) AddEmployee(
	ctx context.Context,
	id, name, phone string,
	year int,
) error {
	// Валидация данных
	if id == "" {
		return errors.New("ID cannot be empty")
	}
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if year < 1900 || year > 2100 {
		return errors.New("invalid year (1900-2100)")
	}

	return s.repo.Add(ctx, Employee{
		ID:    id,
		Name:  name,
		Year:  year,
		Phone: phone,
	})
}

// GetAllEmployees - получение всех сотрудников
func (s *EmployeeService) GetAllEmployees(ctx context.Context) ([]Employee, error) {
	return s.repo.GetAll(ctx)
}

// EmployeeHandler - обработчики HTTP (Transport Layer)
type EmployeeHandler struct {
	service EmployeeUseCase // Зависимость от сервиса
}

// NewEmployeeHandler - конструктор обработчика
func NewEmployeeHandler(service EmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

// AddEmployeeHandler - обработчик добавления
func (h *EmployeeHandler) AddEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Получаем контекст запроса

	// Устанавливаем таймаут 2 секунды
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if r.Method != http.MethodPost { // Проверяем метод
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct { // Структура для парсинга запроса
		ID    string `json:"id"`
		Name  string `json:"name"`
		Year  int    `json:"year"`
		Phone string `json:"phone"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.service.AddEmployee(ctx, req.ID, req.Name, req.Phone, req.Year); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(map[string]string{"status": "created"})
}

// GetAllEmployeesHandler - обработчик получения всех
func (h *EmployeeHandler) GetAllEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Контекст запроса

	if r.Method != http.MethodGet { // Проверка метода
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	employees, err := h.service.GetAllEmployees(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees) // Сериализация в JSON
}

// EmployeeClient - HTTP-клиент для API
type EmployeeClient struct {
	baseURL string       // Базовый URL сервера
	client  *http.Client // HTTP-клиент
}

// NewEmployeeClient - конструктор клиента
func NewEmployeeClient(baseURL string) *EmployeeClient {
	return &EmployeeClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 5 * time.Second, // Таймаут 5 секунд
		},
	}
}

// AddEmployee - метод клиента для добавления
func (c *EmployeeClient) AddEmployee(ctx context.Context, emp Employee) error {
	reqBody, err := json.Marshal(emp) // Сериализация в JSON
	if err != nil {
		return err
	}

	// Создание запроса с контекстом
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/employees/add",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Выполнение запроса
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return nil
}

// GetAllEmployees - метод клиента для получения всех
func (c *EmployeeClient) GetAllEmployees(ctx context.Context) ([]Employee, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+"/employees/all",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var employees []Employee
	if err := json.NewDecoder(resp.Body).Decode(&employees); err != nil {
		return nil, err
	}

	return employees, nil
}

func main() {
	// Инициализация хранилища
	repo := NewInMemoryEmployeeRepository()

	// Инициализация сервиса
	service := NewEmployeeService(repo)

	// Инициализация обработчиков
	handler := NewEmployeeHandler(service)

	// Настройка маршрутов
	http.HandleFunc("/employees/add", handler.AddEmployeeHandler)
	http.HandleFunc("/employees/all", handler.GetAllEmployeesHandler)

	// Запуск сервера в горутине
	go func() {
		fmt.Println("Server starting on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Даем серверу время запуститься
	time.Sleep(100 * time.Millisecond)

	// Пример использования клиента
	client := NewEmployeeClient("http://localhost:8080")
	ctx := context.Background()

	// Добавление сотрудника
	err := client.AddEmployee(ctx, Employee{
		ID:   "1",
		Name: "John Doe",
		Year: 1990,
	})
	if err != nil {
		log.Printf("Add error: %v", err)
	}

	// Получение всех сотрудников
	employees, err := client.GetAllEmployees(ctx)
	if err != nil {
		log.Printf("GetAll error: %v", err)
	} else {
		fmt.Printf("Employees: %+v\n", employees)
	}

	// Ожидание (в реальном приложении может быть select{})
	time.Sleep(1 * time.Second)
}
