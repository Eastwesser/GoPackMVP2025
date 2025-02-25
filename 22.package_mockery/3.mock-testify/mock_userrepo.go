package main

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository - мок-реализация UserRepository
type MockUserRepository struct {
	mock.Mock
}

// FindById - мок-метод
func (m *MockUserRepository) FindById(ctx context.Context, id string) (*User, error) {
	args := m.Called(ctx, id)
	if user, ok := args.Get(0).(*User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}
