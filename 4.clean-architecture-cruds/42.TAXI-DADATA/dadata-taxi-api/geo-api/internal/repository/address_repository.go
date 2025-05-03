// internal/repository/address_repository.go
package repository

import (
	"GoPackMVP2025/42.TAXI-DADATA/dadata-taxi-api/geo-api/internal/entity"
	"context"

	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(ctx context.Context, address *entity.Address) error
	FindByID(ctx context.Context, id uint) (*entity.Address, error)
	FindByUser(ctx context.Context, userID uint) ([]entity.Address, error)
	Delete(ctx context.Context, id uint) error
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepo{db: db}
}

// Реализация методов...
