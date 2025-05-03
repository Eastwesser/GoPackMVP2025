package models

// Модель подписчиков

/*
	Модели базы данных для работы с PostgreSQL:
		subscriber.go: Модель для подписчиков (id, telegram_id, дата подписки).
*/

type Subscriber struct {
	ID         int
	TelegramID int64
	Subscribed bool
}
