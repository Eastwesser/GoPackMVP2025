package repository

import (
	"cleanuser/internal/entity"
	"sync"
)

const repository = "Infrastructure - реализация репозитория"

// --- Infrastructure Layer (Реализация репозитория) ---
// Детали работы с данными (в данном случае in-memory, но обычно это база данных Постгре

//Пояснение:
//Реализует интерфейсы из entity слоя
//Содержит детали работы с данными (в данном случае in-memory хранилище)
//В реальном проекте здесь будет работа с БД, внешними API и т.д.
//Может быть легко заменен на другую реализацию (Postgres, MongoDB и т.д.)

type MemoryUserRepository struct {
	mu     sync.RWMutex
	users  map[int]*entity.User
	nextID int
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users:  make(map[int]*entity.User),
		nextID: 1,
	}
}

func (r *MemoryUserRepository) CreateUser(user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.users[user.ID] = user
	r.nextID++
	return nil
}

func (r *MemoryUserRepository) ReadUser() ([]entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var users []entity.User
	for _, user := range r.users {
		users = append(users, *user)

	}

	return users, nil
}
