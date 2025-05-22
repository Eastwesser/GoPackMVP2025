package controller

import (
	"cleanarch/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func NewEmpRouter(uc *usecase.EmpUseCase) *chi.Mux {
	r := chi.NewRouter()
	handler := NewEmpHandler(uc)

	// Используем RESTful стиль путей
	r.Post("/employees", handler.AddEmployee)    // POST /employees
	r.Get("/employees", handler.GetAllEmployees) // GET /employees

	return r
}
