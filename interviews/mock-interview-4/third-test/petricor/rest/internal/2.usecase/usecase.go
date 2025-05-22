package __usecase

import (
	"context"
	"errors"
	"internal/__entity"
	"internal/__repository"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *__entity.User) error
	GetAllUsers(ctx context.Context) ([]__entity.User, error)
}

type UserUsecase struct {
	repo __repository.IUserRepository
}

func NewUserUsecase(repo __repository.IUserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) RegisterUser(ctx context.Context, user *__entity.User) error {
	if user.Age < 18 {
		return errors.New("user must be adult") // Упрощённо, можно свою ошибку
	}
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsecase) GetAllUsers(ctx context.Context) ([]__entity.User, error) {
	return uc.repo.GetAllUsers(ctx)
}
