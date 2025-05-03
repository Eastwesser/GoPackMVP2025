package usecase

import "andrius-cleanarch-map/internal/entity"

// interfaces and business logic

type UserRepository interface {
	Save(user entity.User)
	FindAll() []entity.User
	FindByID(id string) (entity.User, bool)
}

type UserInteractor struct {
	repo UserRepository
}

func NewUserInteractor(r UserRepository) *UserInteractor {
	return &UserInteractor{repo: r}
}

func (u *UserInteractor) CreateUser(id, name string) {
	user := entity.User{ID: id, Name: name}
	u.repo.Save(user)
}

func (u *UserInteractor) GetAll() []entity.User {
	return u.repo.FindAll()
}

func (u *UserInteractor) GetByID(id string) (entity.User, bool) {
	return u.repo.FindByID(id)
}
