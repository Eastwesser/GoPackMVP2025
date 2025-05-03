package main

import (
	"errors"
	"fmt"
)

type Entity struct {
	ID    string
	Name  string
	Value int
}

type Storage interface {
	Create(entity Entity) error
	Read(id string) (Entity, error)
	Update(id string, entity Entity) error
	Delete(id string) error
}

type InMemoryStorage struct {
	entities map[string]Entity
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		entities: make(map[string]Entity),
	}
}

func (s *InMemoryStorage) Create(entity Entity) error {
	if _, exists := s.entities[entity.ID]; exists {
		return errors.New("entity already exists")
	}
	s.entities[entity.ID] = entity
	return nil
}

func (s *InMemoryStorage) Read(id string) (Entity, error) {
	entity, exists := s.entities[id]
	if !exists {
		return Entity{}, errors.New("entity not found")
	}
	return entity, nil
}

func (s *InMemoryStorage) Update(id string, entity Entity) error {
	if _, exists := s.entities[id]; !exists {
		return errors.New("entity not found")
	}
	s.entities[id] = entity
	return nil
}

func (s *InMemoryStorage) Delete(id string) error {
	if _, exists := s.entities[id]; !exists {
		return errors.New("entity not found")
	}
	delete(s.entities, id)
	return nil
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) createEntity(id, name string, value int) error {
	return s.storage.Create(Entity{ID: id, Name: name, Value: value})
}

func (s *Service) readEntity(id string) (Entity, error) {
	return s.storage.Read(id)
}

func (s *Service) updateEntity(id, name string, value int) error {
	return s.storage.Update(id, Entity{ID: id, Name: name, Value: value})
}

func (s *Service) deleteEntity(id string) error {
	return s.storage.Delete(id)
}

func main() {
	fmt.Println("CRUD Operations Example")

	storage := NewInMemoryStorage()
	service := NewService(storage)

	// Create
	err := service.createEntity("1", "First", 100)
	if err != nil {
		fmt.Println("Create error:", err)
	}

	// Read
	entity, err := service.readEntity("1")
	if err != nil {
		fmt.Println("Read error:", err)
	} else {
		fmt.Printf("Read entity: %+v\n", entity)
	}

	// Update
	err = service.updateEntity("1", "Updated", 200)
	if err != nil {
		fmt.Println("Update error:", err)
	}

	// Read after update
	entity, err = service.readEntity("1")
	if err != nil {
		fmt.Println("Read error:", err)
	} else {
		fmt.Printf("Read after update: %+v\n", entity)
	}

	// Delete
	err = service.deleteEntity("1")
	if err != nil {
		fmt.Println("Delete error:", err)
	}

	// Try to read after delete
	_, err = service.readEntity("1")
	if err != nil {
		fmt.Println("Read after delete error:", err)
	}
}
