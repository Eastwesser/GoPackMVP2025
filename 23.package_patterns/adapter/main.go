package main

import (
	"context"
	"errors"
)

type User struct {
	Name  string
	Email string
	Id    string
}

// UserRepositoryOld — это старый интерфейс, который нужно адаптировать.
type UserRepositoryOld interface {
	FindById(context.Context, string) (*User, error)
}

// UserService — это новый интерфейс, который ожидает клиент.
type UserService interface {
	GetUser(context.Context, string) (*User, error)
}

// UserRepositoryAdapter — это адаптер, который делает UserRepository совместимым с UserService.
type UserRepositoryAdapter struct {
	repo UserRepositoryOld // Адаптируемый объект (UserRepository)
}

// GetUser — метод, который реализует интерфейс UserService.
func (ura *UserRepositoryAdapter) GetUser(ctx context.Context, id string) (*User, error) {
	user, err := ura.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// =====================================================================================================================

// MockUserRepository — моковая реализация UserRepository для тестирования.
type MockUserRepository struct{}

func (mur *MockUserRepository) FindById(ctx context.Context, id string) (*User, error) {
	if id == "1" {
		return &User{Id: "1", Name: "Denis", Email: "denis@example.com"}, nil
	}
	return nil, errors.New("user not found")
}

// =====================================================================================================================

// Пример использования
func main() {
	// Создаем моковый UserRepository
	mockRepo := &MockUserRepository{}

	// Создаем адаптер
	adapter := &UserRepositoryAdapter{
		repo: mockRepo,
	}

	// Используем UserService через адаптер
	user, err := adapter.GetUser(context.Background(), "1")
	if err != nil {
		println("Error:", err.Error())
		return
	}

	println("User:", user.Name, user.Email)
}
