package controller

import (
	"cleanarch/internal/domain"
	"cleanarch/internal/usecase"
	"encoding/json"
	"net/http"
)

type EmpHandler struct {
	uc *usecase.EmpUseCase
}

func NewEmpHandler(uc *usecase.EmpUseCase) *EmpHandler {
	return &EmpHandler{uc}
}

func (h *EmpHandler) AddEmployee(w http.ResponseWriter, r *http.Request) {
	var emp domain.Emp

	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.uc.AddEmp(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func (h *EmpHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.uc.GetAllEmp()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
