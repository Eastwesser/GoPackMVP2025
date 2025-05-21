package repository

import (
	"cleanarch/internal/domain"
	"errors"
	"sync"
)

type empRepo struct {
	mu        sync.RWMutex
	employees map[int]*domain.Emp
}

func NewEmpRepo() *empRepo {
	return &empRepo{
		employees: make(map[int]*domain.Emp),
	}
}

func (r *empRepo) Add(emp *domain.Emp) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exist := r.employees[emp.ID]; exist {
		return errors.New("employee already exists")
	}

	r.employees[emp.ID] = emp

	return nil
}

func (r *empRepo) GetAll() ([]domain.Emp, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]domain.Emp, 0, len(r.employees))
	for _, emp := range r.employees {
		result = append(result, *emp)
	}

	return result, nil
}
