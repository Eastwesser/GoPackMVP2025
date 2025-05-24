package entity

// --- Domain Layer (Сущности) ---
const entity = "Domain - ядро без зависимостей"

// User это чистая бизнес-логика (ядро системы), не зависит ни от чего внешнего
// Содержит только структуры данных и интерфейсы
type User struct {
	ID       int
	Name     string
	Age      int
	Nickname string
	Phone    string
	Email    string
}
