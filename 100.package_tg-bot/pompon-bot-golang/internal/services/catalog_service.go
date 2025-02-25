package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"pompon-bot-golang/internal/models"
)

// CatalogService предоставляет методы для работы с каталогом товаров.
type CatalogService struct {
	db *pgxpool.Pool
}

// NewCatalogService создает новый экземпляр CatalogService.
func NewCatalogService(db *pgxpool.Pool) *CatalogService {
	return &CatalogService{db: db}
}

// GetCategories возвращает список категорий товаров.
func (s *CatalogService) GetCategories(ctx context.Context) ([]string, error) {
	rows, err := s.db.Query(ctx, "SELECT name FROM categories")
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, name)
	}

	return categories, nil
}

// GetProductsByCategory возвращает список товаров по категории.
func (s *CatalogService) GetProductsByCategory(ctx context.Context, category string) ([]models.Product, error) {
	rows, err := s.db.Query(ctx, "SELECT id, name, description, price FROM products WHERE category = $1", category)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}
