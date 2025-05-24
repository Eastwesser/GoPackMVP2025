package usecase

import (
	"cleanuser/internal/entity"
	"errors"
)

const usecase = "Usecase - бизнес-правила"

// IUserRepository
// Интерфейсы репозиториев определяются здесь, но реализуются во внешних слоях
// --- Use Case Layer (Сценарии) ---
// Бизнес-правила приложения
type IUserRepository interface {
	CreateUser(user *entity.User) error
	ReadUser() ([]entity.User, error)
}

type UserUseCase struct {
	repo IUserRepository
}

func NewUserUseCase(repo IUserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) CreateUser(user *entity.User) error {

	if user.Age < 18 {
		return errors.New("user must be at least 18 years old")
	}

	return uc.repo.CreateUser(user)
}

func (uc *UserUseCase) ReadUser() ([]entity.User, error) {
	return uc.repo.ReadUser()
}
