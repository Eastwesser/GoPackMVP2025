package infrastructure

import (
	"andrius-cleanarch-map/internal/entity"
)

// MAP REALISATION

type InMemoryUserRepo struct {
	data map[string]entity.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{data: make(map[string]entity.User)}
}

func (r *InMemoryUserRepo) Save(user entity.User) {
	r.data[user.ID] = user
}

func (r *InMemoryUserRepo) FindAll() []entity.User {
	users := []entity.User{}
	for _, u := range r.data {
		users = append(users, u)
	}
	return users
}

func (r *InMemoryUserRepo) FindByID(id string) (entity.User, bool) {
	user, ok := r.data[id]
	return user, ok
}
