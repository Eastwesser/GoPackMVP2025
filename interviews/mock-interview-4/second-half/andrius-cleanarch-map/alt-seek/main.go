package main

import (
	"fmt"
)

// Domain layer
type Entity struct {
	ID   string
	Name string
}

// Repository interface (port)
type Repository interface {
	Create(id string, entity Entity) error
	Get(id string) (Entity, error)
	Update(id string, entity Entity) error
	Delete(id string) error
}

// UseCase interface (port)
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
	return &InMemoryRepository{
		storage: make(map[string]Entity),
	}
}

func (r *InMemoryRepository) Create(id string, entity Entity) error {
	if _, exists := r.storage[id]; exists {
		return fmt.Errorf("entity with id %s already exists", id)
	}
	r.storage[id] = entity
	return nil
}

func (r *InMemoryRepository) Get(id string) (Entity, error) {
	entity, exists := r.storage[id]
	if !exists {
		return Entity{}, fmt.Errorf("entity with id %s not found", id)
	}
	return entity, nil
}

func (r *InMemoryRepository) Update(id string, entity Entity) error {
	if _, exists := r.storage[id]; !exists {
		return fmt.Errorf("entity with id %s not found", id)
	}
	r.storage[id] = entity
	return nil
}

func (r *InMemoryRepository) Delete(id string) error {
	if _, exists := r.storage[id]; !exists {
		return fmt.Errorf("entity with id %s not found", id)
	}
	delete(r.storage, id)
	return nil
}

// Application layer
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

// Presentation layer (main)
func main() {
	repo := NewInMemoryRepository()
	useCase := NewEntityUseCase(repo)

	// Пример использования
	err := useCase.CreateEntity("1", "First Entity")
	if err != nil {
		fmt.Println("Error:", err)
	}

	entity, err := useCase.GetEntity("1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Got entity: %+v\n", entity)
	}

	err = useCase.UpdateEntity("1", "Updated Entity")
	if err != nil {
		fmt.Println("Error:", err)
	}

	entity, err = useCase.GetEntity("1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Got updated entity: %+v\n", entity)
	}

	err = useCase.DeleteEntity("1")
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = useCase.GetEntity("1")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
