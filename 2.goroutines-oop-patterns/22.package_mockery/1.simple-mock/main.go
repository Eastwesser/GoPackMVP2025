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

type UserRepository interface {
	FindById(context.Context, string) (*User, error)
}

type UserRepositoryImpl struct {
	repo UserRepository
}

func (uri *UserRepositoryImpl) GetUser(ctx context.Context, id string) (*User, error) {

	user, err := uri.repo.FindById(ctx, id)

	// в моках этот метод FindById четко захардкожен,
	// мы бы точно знали, что там ("1": {Id: "1", Name: "Denis", Email: "denis@example.com"},)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// МОК ОБЪЕКТ (имитации)

type MockUserRepository struct {
	mockData map[string]*User
}

func (m *MockUserRepository) FindById(ctx context.Context, id string) (*User, error) {

	if user, ok := m.mockData[id]; ok {
		return user, nil
	}

	return nil, errors.New("user not found")
}
