package repository

import (
	"cleanuser/internal/domain"
	"errors"
	"sync"
)

// --- Infrastructure Layer (Реализация репозитория) ---
// Детали работы с данными (в данном случае in-memory, но обычно это база данных Постгре

//Пояснение:
//Реализует интерфейсы из domain слоя
//Содержит детали работы с данными (в данном случае in-memory хранилище)
//В реальном проекте здесь будет работа с БД, внешними API и т.д.
//Может быть легко заменен на другую реализацию (Postgres, MongoDB и т.д.)

type MemoryUserRepository struct {
	mu     sync.RWMutex
	users  map[int]*domain.User
	nextID int
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users:  make(map[int]*domain.User),
		nextID: 1,
	}
}

func (r *MemoryUserRepository) Create(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.users[user.ID] = user
	r.nextID++
	return nil
}

func (r *MemoryUserRepository) GetByID(id int) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *MemoryUserRepository) Update(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}

func (r *MemoryUserRepository) List() ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}
	return users, nil
}
