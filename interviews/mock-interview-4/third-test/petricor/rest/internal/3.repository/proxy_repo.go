package __repository

import (
	"context"
	"internal/__entity"
	"time"
)

type CachedUserRepo struct {
	repo     IUserRepository
	cache    map[string][]__entity.User // Простейший in-memory кэш
	cacheTTL time.Duration
}

func NewCachedUserRepo(repo IUserRepository, ttl time.Duration) *CachedUserRepo {
	return &CachedUserRepo{
		repo:     repo,
		cache:    make(map[string][]__entity.User),
		cacheTTL: ttl,
	}
}

func (r *CachedUserRepo) CreateUser(ctx context.Context, user *__entity.User) error {
	return r.repo.CreateUser(ctx, user) // Проксируем без кэширования
}

func (r *CachedUserRepo) GetAllUsers(ctx context.Context) ([]__entity.User, error) {
	cacheKey := "all_users"

	// Проверяем кэш
	if users, ok := r.cache[cacheKey]; ok {
		return users, nil
	}

	// Если нет в кэше — запрашиваем у репозитория
	users, err := r.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	// Сохраняем в кэш
	r.cache[cacheKey] = users
	go func() {
		time.Sleep(r.cacheTTL)
		delete(r.cache, cacheKey) // Очистка по TTL
	}()

	return users, nil
}
