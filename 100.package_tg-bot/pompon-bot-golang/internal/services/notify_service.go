package services

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NotifyService предоставляет методы для рассылки уведомлений.
type NotifyService struct {
	db  *pgxpool.Pool
	bot *tgbotapi.BotAPI
}

// NewNotifyService создает новый экземпляр NotifyService.
func NewNotifyService(db *pgxpool.Pool, bot *tgbotapi.BotAPI) *NotifyService {
	return &NotifyService{db: db, bot: bot}
}

// NotifyAllSubscribers отправляет сообщение всем подписчикам.
func (s *NotifyService) NotifyAllSubscribers(ctx context.Context, message string) error {
	rows, err := s.db.Query(ctx, "SELECT telegram_id FROM subscribers")
	if err != nil {
		return fmt.Errorf("failed to get subscribers: %w", err)
	}
	defer rows.Close()

	var subscribers []int64
	for rows.Next() {
		var telegramID int64
		if err := rows.Scan(&telegramID); err != nil {
			return fmt.Errorf("failed to scan subscriber: %w", err)
		}
		subscribers = append(subscribers, telegramID)
	}

	// Отправляем сообщение каждому подписчику.
	for _, telegramID := range subscribers {
		msg := tgbotapi.NewMessage(telegramID, message)
		if _, err := s.bot.Send(msg); err != nil {
			return fmt.Errorf("failed to send message to subscriber %d: %w", telegramID, err)
		}
	}

	return nil
}
