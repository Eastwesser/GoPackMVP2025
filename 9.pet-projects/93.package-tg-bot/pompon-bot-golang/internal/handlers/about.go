package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleStart processes the /start command
func HandleStart(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to the POMPON store! 🎁\nPlease select an action:")
	msg.ReplyMarkup = MainMenuKeyboard()
	bot.Send(msg)
}

// HandleAbout processes the /about command
func HandleAbout(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "We are the POMPON store! 🎁\nHere you will find the best gift boxes, cards, and gift wraps.")
	bot.Send(msg)
}

// MainMenuKeyboard creates a keyboard for the main menu
func MainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🔹 About Us"),
			tgbotapi.NewKeyboardButton("📦 Catalog"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🛒 Order"),
			tgbotapi.NewKeyboardButton("🔔 Subscription"),
		),
	)
}
