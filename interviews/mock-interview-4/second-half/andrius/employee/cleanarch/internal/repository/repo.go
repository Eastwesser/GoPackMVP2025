package repository

import (
	"cleanarch/internal/entity"
	"errors"
	"sync"
)

// EmpRepo - in-memory реализация хранилища сотрудников
type IEmpRepo struct {
	mu        sync.RWMutex        // Для потокобезопасности
	employees map[int]*entity.Emp // Хранение данных в map (ключ - ID)
}

// NewEmpRepo создает новый репозиторий
func NewEmpRepo() *IEmpRepo {
	return &IEmpRepo{
		employees: make(map[int]*entity.Emp),
	}
}

// AddEmp добавляет нового сотрудника
func (r *IEmpRepo) AddEmp(emp *entity.Emp) error {
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

// GetAllEmp возвращает всех сотрудников
func (r *IEmpRepo) GetAllEmp() ([]entity.Emp, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Конвертируем map в slice
	result := make([]entity.Emp, 0, len(r.employees))
	for _, emp := range r.employees {
		result = append(result, *emp)
	}

	return result, nil
}
