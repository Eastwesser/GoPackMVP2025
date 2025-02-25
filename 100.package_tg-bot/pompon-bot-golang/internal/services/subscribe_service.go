package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SubscribeService предоставляет методы для работы с подписками.
type SubscribeService struct {
	db *pgxpool.Pool
}

// NewSubscribeService создает новый экземпляр SubscribeService.
func NewSubscribeService(db *pgxpool.Pool) *SubscribeService {
	return &SubscribeService{db: db}
}

// AddSubscriber добавляет подписчика, если его еще нет в базе.
func (s *SubscribeService) AddSubscriber(ctx context.Context, telegramID int64) error {
	// Проверяем, есть ли уже подписчик
	var existingID int64
	err := s.db.QueryRow(ctx, "SELECT telegram_id FROM subscribers WHERE telegram_id = $1", telegramID).Scan(&existingID)
	if err == nil {
		// Если подписчик уже есть, ничего не делаем
		return nil
	}

	// Если подписчика нет, добавляем его
	_, err = s.db.Exec(ctx, "INSERT INTO subscribers (telegram_id) VALUES ($1) ON CONFLICT DO NOTHING", telegramID)
	if err != nil {
		return fmt.Errorf("failed to add subscriber: %w", err)
	}

	return nil
}

// RemoveSubscriber удаляет подписчика.
func (s *SubscribeService) RemoveSubscriber(ctx context.Context, telegramID int64) error {
	_, err := s.db.Exec(ctx, "DELETE FROM subscribers WHERE telegram_id = $1", telegramID)
	if err != nil {
		return fmt.Errorf("failed to remove subscriber: %w", err)
	}

	return nil
}

// ListSubscribers возвращает список всех подписчиков.
func (s *SubscribeService) ListSubscribers(ctx context.Context) ([]int64, error) {
	rows, err := s.db.Query(ctx, "SELECT telegram_id FROM subscribers")
	if err != nil {
		return nil, fmt.Errorf("failed to get subscribers: %w", err)
	}
	defer rows.Close()

	var subscribers []int64
	for rows.Next() {
		var telegramID int64
		if err := rows.Scan(&telegramID); err != nil {
			return nil, fmt.Errorf("failed to scan subscriber: %w", err)
		}
		subscribers = append(subscribers, telegramID)
	}

	return subscribers, nil
}
