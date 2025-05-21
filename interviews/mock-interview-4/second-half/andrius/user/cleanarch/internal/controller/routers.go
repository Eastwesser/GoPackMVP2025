package controller

import (
	"cleanarch/internal/usecase"
)

func NewEmpRouter(uc *usecase.EmpUseCase) *chi.Mux {
	r := chi.NewRouter()
	handler := NewEmpHandler(uc)

	r.Post("/add", handler.AddEmp)
	r.Get("/getall", handler.GetAllEmp)

	return r
}
