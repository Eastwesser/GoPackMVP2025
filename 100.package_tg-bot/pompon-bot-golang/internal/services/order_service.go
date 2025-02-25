package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"pompon-bot-golang/internal/models"
)

// OrderService предоставляет методы для работы с заказами.
type OrderService struct {
	db *pgxpool.Pool
}

// NewOrderService создает новый экземпляр OrderService.
func NewOrderService(db *pgxpool.Pool) *OrderService {
	return &OrderService{db: db}
}

// GetCategories возвращает список категорий товаров.
func (s *OrderService) GetCategories(ctx context.Context) ([]string, error) {
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

// CreateOrder создает новый заказ, привязывая его к telegram_id.
func (s *OrderService) CreateOrder(ctx context.Context, telegramID int64, order models.Order) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Вставляем заказ, привязывая его к telegram_id
	_, err = tx.Exec(ctx, "INSERT INTO orders (user_id, product_id, quantity, status) VALUES ($1, $2, $3, $4)",
		telegramID, order.Product.ID, order.Quantity, order.Status)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// GetOrdersByUser возвращает заказы пользователя, используя telegram_id.
func (s *OrderService) GetOrdersByUser(ctx context.Context, telegramID int64) ([]models.Order, error) {
	rows, err := s.db.Query(ctx, "SELECT id, product_id, quantity, status FROM orders WHERE user_id = $1", telegramID)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.Product.ID, &order.Quantity, &order.Status); err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}
