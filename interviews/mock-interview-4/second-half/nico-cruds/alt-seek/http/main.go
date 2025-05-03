package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Domain layer
type Entity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Repository interface
type Repository interface {
	Create(id string, entity Entity) error
	Get(id string) (Entity, error)
	Update(id string, entity Entity) error
	Delete(id string) error
}

// UseCase interface
type UseCase interface {
	CreateEntity(id string, name string) error
	GetEntity(id string) (Entity, error)
	UpdateEntity(id string, name string) error
	DeleteEntity(id string) error
}

// Infrastructure layer
type InMemoryRepository struct {
	storage map[string]Entity
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{storage: make(map[string]Entity)}
}

func (r *InMemoryRepository) Create(id string, entity Entity) error {
	if _, exists := r.storage[id]; exists {
		return fmt.Errorf("entity already exists")
	}
	r.storage[id] = entity
	return nil
}

func (r *InMemoryRepository) Get(id string) (Entity, error) {
	entity, exists := r.storage[id]
	if !exists {
		return Entity{}, fmt.Errorf("entity not found")
	}
	return entity, nil
}

func (r *InMemoryRepository) Update(id string, entity Entity) error {
	if _, exists := r.storage[id]; !exists {
		return fmt.Errorf("entity not found")
	}
	r.storage[id] = entity
	return nil
}

func (r *InMemoryRepository) Delete(id string) error {
	if _, exists := r.storage[id]; !exists {
		return fmt.Errorf("entity not found")
	}
	delete(r.storage, id)
	return nil
}

// UseCase
type EntityUseCase struct {
	repo Repository
}

func NewEntityUseCase(repo Repository) *EntityUseCase {
	return &EntityUseCase{repo: repo}
}

func (uc *EntityUseCase) CreateEntity(id string, name string) error {
	return uc.repo.Create(id, Entity{ID: id, Name: name})
}

func (uc *EntityUseCase) GetEntity(id string) (Entity, error) {
	return uc.repo.Get(id)
}

func (uc *EntityUseCase) UpdateEntity(id string, name string) error {
	entity, err := uc.repo.Get(id)
	if err != nil {
		return err
	}
	entity.Name = name
	return uc.repo.Update(id, entity)
}

func (uc *EntityUseCase) DeleteEntity(id string) error {
	return uc.repo.Delete(id)
}

// **НОВОЕ: HTTP-хэндлеры (Transport Layer)**
type HTTPHandler struct {
	useCase UseCase
}

func NewHTTPHandler(useCase UseCase) *HTTPHandler {
	return &HTTPHandler{useCase: useCase}
}

func (h *HTTPHandler) CreateEntityHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.useCase.CreateEntity(req.ID, req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "created"})
}

func (h *HTTPHandler) GetEntityHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	entity, err := h.useCase.GetEntity(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity)
}

func (h *HTTPHandler) UpdateEntityHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.useCase.UpdateEntity(req.ID, req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

func (h *HTTPHandler) DeleteEntityHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	if err := h.useCase.DeleteEntity(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}

func main() {
	// Инициализация слоев
	repo := NewInMemoryRepository()
	useCase := NewEntityUseCase(repo)
	handler := NewHTTPHandler(useCase)

	// Роутинг HTTP
	http.HandleFunc("/entity/create", handler.CreateEntityHandler)
	http.HandleFunc("/entity/get", handler.GetEntityHandler)
	http.HandleFunc("/entity/update", handler.UpdateEntityHandler)
	http.HandleFunc("/entity/delete", handler.DeleteEntityHandler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
