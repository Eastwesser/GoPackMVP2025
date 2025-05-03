package handlers

import (
	"context"
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"pompon-bot-golang/internal/services"
	"strings"
)

// HandleSubscribe обрабатывает команду /subscribe
func HandleSubscribe(subscribeService *services.SubscribeService, bot *tgbotapi.BotAPI, chatID int64) {
	ctx := context.Background()

	// Добавляем пользователя в таблицу subscribers
	err := subscribeService.AddSubscriber(ctx, chatID)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Ошибка при подписке. Попробуйте позже.")
		bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(chatID, "Вы успешно подписались на рассылку! 📩")
	bot.Send(msg)
}

// HandleUnsubscribe обрабатывает отписку
func HandleUnsubscribe(subscribeService *services.SubscribeService, bot *tgbotapi.BotAPI, chatID int64) {
	ctx := context.Background()
	err := subscribeService.RemoveSubscriber(ctx, chatID)

	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Ошибка при отписке. Попробуйте позже.")
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatID, "Вы успешно отписались от рассылки. ❌")
	bot.Send(msg)
}

func notifySubscribers(bot *tgbotapi.BotAPI, db *sql.DB, message string) {
	rows, err := db.Query("SELECT telegram_id FROM subscribers")
	if err != nil {
		log.Println("Failed to get subscribers:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var chatID int64
		if err := rows.Scan(&chatID); err != nil {
			log.Println("Failed to scan subscriber:", err)
			continue
		}

		_, err := bot.Send(tgbotapi.NewMessage(chatID, message))
		if err != nil {
			log.Printf("Failed to send message to subscriber %d: %v\n", chatID, err)
			// Если не удается отправить сообщение, значит, пользователь больше не существует
			if strings.Contains(err.Error(), "chat not found") {
				// Удаляем из таблицы подписчика, если не удалось отправить сообщение
				_, err := db.Exec("DELETE FROM subscribers WHERE telegram_id = $1", chatID)
				if err != nil {
					log.Println("Failed to delete subscriber:", err)
				} else {
					log.Println("Subscriber deleted:", chatID)
				}
			}
		}
	}

	// Проверяем на ошибки после завершения обхода rows
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
	}
}
