package __repository

import (
	"context"
	"errors"
	"internal/__entity"
)

// Mock-репозиторий (в реальности тут будет БД)
type MockUserRepo struct{}

func (r *MockUserRepo) CreateUser(ctx context.Context, user *__entity.User) error {
	// Проверка возраста (можно вынести в usecase)
	if user.Age < 18 {
		return errors.New("age must be >= 18")
	}
	// Здесь была бы проверка на существующего пользователя в БД
	panic("implement me: check if user exists by email")
	return nil
}

func (r *MockUserRepo) GetAllUsers(ctx context.Context) ([]__entity.User, error) {
	panic("implement me: fetch users from DB")
}
