package keyboards

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Генерация клавиатур

/*
	Генерация клавиатур для удобного взаимодействия:
		Кнопки для выбора категории.
		Подтверждение заказа.
		Уведомления о подписке.
*/

// CatalogKeyboard создает inline-клавиатуру для выбора категорий
func CatalogKeyboard(categories []string) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	for _, category := range categories {
		button := tgbotapi.NewInlineKeyboardButtonData(category, fmt.Sprintf("category_%s", category))
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(button))
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

// OrderKeyboard создает клавиатуру для подтверждения заказа
func OrderKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Подтвердить заказ ✅"),
			tgbotapi.NewKeyboardButton("Отменить заказ ❌"),
		),
	)
}

// SubscribeKeyboard создает клавиатуру для подписки
func SubscribeKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Подписаться 📩"),
			tgbotapi.NewKeyboardButton("Отписаться ❌"),
		),
	)
}
