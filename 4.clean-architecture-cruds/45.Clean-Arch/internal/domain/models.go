package domain

// --- Domain Layer (Сущности) ---
// Это чистая бизнес-логика (ядро системы), не зависит ни от чего внешнего
// Содержит только структуры данных и интерфейсы
type User struct {
	ID       int
	Name     string
	Age      int
	Nickname string
	Phone    string
	Email    string
}

// UserRepository определен в domain, но реализуется в infrastructure
// Интерфейсы репозиториев определяются здесь, но реализуются во внешних слоях
// --- Use Case Layer (Сценарии) ---
// Бизнес-правила приложения
type UserRepository interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	Update(user *User) error
	Delete(id int) error
	List() ([]User, error)
}
