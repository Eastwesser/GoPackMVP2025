package controller

import (
	"cleanarch/internal/entity"
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
	var emp entity.Emp

	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.uc.Add(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(emp)
	if err != nil {
		return
	}
}

func (h *EmpHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.uc.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(employees)
	if err != nil {
		return
	}
}
