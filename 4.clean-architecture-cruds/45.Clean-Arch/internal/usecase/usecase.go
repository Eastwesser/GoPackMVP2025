package usecase

import (
	"cleanuser/internal/domain"
	"errors"
)

const usecase = "Usecase - бизнес-правила"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Create(user *domain.User) error {

	if user.Age < 18 {
		return errors.New("user must be at least 18 years old")
	}

	return uc.repo.Create(user)
}

func (uc *UserUseCase) GetByID(id int) (*domain.User, error) {
	return uc.repo.Get(id)
}

func (uc *UserUseCase) Update(user *domain.User) error {
	existing, err := uc.repo.Get(user.ID)
	if err != nil {
		return err
	}

	// Business rule: email cannot be changed
	user.Email = existing.Email

	return uc.repo.Update(user)
}

func (uc *UserUseCase) Delete(id int) error {
	return uc.repo.Delete(id)
}

//
//func (uc *UserUseCase) List() ([]domain.User, error) {
//	return uc.repo.List()
//}
