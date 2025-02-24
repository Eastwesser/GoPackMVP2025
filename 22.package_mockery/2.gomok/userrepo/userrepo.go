package userrepo

import (
	"context"
)

// User - структура пользователя
type User struct {
	Name  string
	Email string
	Id    string
}

// UserRepository - интерфейс репозитория пользователей

//go:generate mockgen -source=userrepo.go -destination=mock_userrepo.go -package=userrepo

type UserRepository interface {
	FindById(ctx context.Context, id string) (*User, error)
}

// UserRepositoryImpl - реализация UserRepository
type UserRepositoryImpl struct {
	repo UserRepository
}

// GetUser - метод получения пользователя по ID
func (uri *UserRepositoryImpl) GetUser(ctx context.Context, id string) (*User, error) {
	return uri.repo.FindById(ctx, id)
}
