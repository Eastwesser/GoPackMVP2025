package usecase

import "cleanarch/internal/entity"

type IEmpRepo interface {
	Add(emp *entity.Emp) error
	GetAll() ([]entity.Emp, error)
}

type EmpUseCase struct {
	repo IEmpRepo
}

// NewEmpUseCase constructor
func NewEmpUseCase(repo IEmpRepo) *EmpUseCase {
	return &EmpUseCase{repo: repo}
}

func (uc *EmpUseCase) Add(emp *entity.Emp) error {
	return uc.repo.Add(emp)
}

func (uc *EmpUseCase) GetAll() ([]entity.Emp, error) {
	return uc.repo.GetAll()
}
