package repository

import (
	"cleanarch/internal/entity"
	"sync"
	"time"
)

type cachedEmployee struct {
	data   []entity.Emp
	expiry time.Time
}

type CacheProxyRepo struct {
	repo     IEmpRepo // Используем интерфейс вместо конкретной реализации
	cache    *cachedEmployee
	mu       sync.RWMutex
	cacheTTL time.Duration
}

// NewCacheProxyRepo создает прокси с кэшированием
func NewCacheProxyRepo(repo *IEmpRepo, cacheTTL time.Duration) *CacheProxyRepo {
	return &CacheProxyRepo{
		repo: *repo,
		cache: &cachedEmployee{
			data:   make([]entity.Emp, 0),
			expiry: time.Now().Add(cacheTTL),
		},
		cacheTTL: cacheTTL,
	}
}

func (r *CacheProxyRepo) Add(emp *entity.Emp) error {
	// Добавляем через реальный репозиторий
	if err := r.repo.AddEmp(emp); err != nil {
		return err
	}

	// Инвалидируем кэш
	r.mu.Lock()
	r.cache = nil
	r.mu.Unlock()

	return nil
}

func (r *CacheProxyRepo) GetAll() ([]entity.Emp, error) {
	r.mu.RLock()
	cached := r.cache
	r.mu.RUnlock()

	// Если кэш валиден - возвращаем его
	if cached != nil && time.Now().Before(cached.expiry) {
		return cached.data, nil
	}

	// Получаем данные из основного репозитория
	employees, err := r.repo.GetAllEmp()
	if err != nil {
		return nil, err
	}

	// Обновляем кэш
	r.mu.Lock()
	r.cache = &cachedEmployee{
		data:   employees,
		expiry: time.Now().Add(r.cacheTTL),
	}
	r.mu.Unlock()

	return employees, nil
}
