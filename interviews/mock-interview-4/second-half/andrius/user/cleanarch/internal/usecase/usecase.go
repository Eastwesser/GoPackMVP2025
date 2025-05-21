package usecase

import "cleanarch/internal/domain"

type EmpUseCase struct {
	repo domain.IEmpRepo
}

func NewEmpUseCase(repo domain.IEmpRepo) *EmpUseCase {
	return &EmpUseCase{repo: repo}
}

func (uc *EmpUseCase) AddEmp(emp *domain.Emp) error {
	return uc.repo.Add(emp)
}

func (uc *EmpUseCase) GetAllEmp() ([]domain.Emp, error) {
	return uc.repo.GetAll()
}
