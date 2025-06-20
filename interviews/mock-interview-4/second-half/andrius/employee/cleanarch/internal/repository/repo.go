package repository

import (
	"cleanarch/internal/entity"
	"errors"
	"sync"
)

// EmpRepo - in-memory реализация хранилища сотрудников
type EmpRepo struct {
	mu        sync.RWMutex        // Для потокобезопасности
	employees map[int]*entity.Emp // Хранение данных в map (ключ - ID)
}

// NewEmpRepo создает новый репозиторий
func NewEmpRepo() *EmpRepo {
	return &EmpRepo{
		employees: make(map[int]*entity.Emp),
	}
}

// Add добавляет нового сотрудника
func (r *EmpRepo) Add(emp *entity.Emp) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Проверка на дубликат ID
	if _, exist := r.employees[emp.ID]; exist {
		return errors.New("employee already exists")
	}

	// Сохраняем сотрудника
	r.employees[emp.ID] = emp
	return nil
}

// GetAll возвращает slice всех сотрудников
func (r *EmpRepo) GetAll() ([]entity.Emp, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Конвертируем map в slice
	result := make([]entity.Emp, 0, len(r.employees))
	for _, emp := range r.employees {
		result = append(result, *emp)
	}

	return result, nil
}
