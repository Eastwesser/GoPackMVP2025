package handlers

import (
	"context"
	"fmt"
	"log"
	"pompon-bot-golang/internal/keyboards"
	"pompon-bot-golang/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// HandleCatalog обрабатывает команду /catalog
func HandleCatalog(catalogService *services.CatalogService, bot *tgbotapi.BotAPI, chatID int64) {
	categories, err := catalogService.GetCategories(context.Background())
	if err != nil {
		log.Printf("Ошибка при получении категорий: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при загрузке категорий.")
		bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(chatID, "Выберите категорию:")
	msg.ReplyMarkup = keyboards.CatalogKeyboard(categories)
	bot.Send(msg)
}

// HandleCallbackQuery обрабатывает callback-запросы
func HandleCallbackQuery(db *pgxpool.Pool, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	data := callbackQuery.Data
	log.Printf("Полученные данные callback: %s", data)

	catalogService := services.NewCatalogService(db)
	if products, err := catalogService.GetProductsByCategory(context.Background(), data); err == nil {
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Товары в категории:")
		for _, product := range products {
			msg.Text += fmt.Sprintf("\n- %s: %s (%.2f руб.)", product.Name, product.Description, product.Price)
		}
		bot.Send(msg)
	} else {
		log.Printf("Ошибка при получении товаров: %v", err)
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Произошла ошибка при загрузке товаров.")
		bot.Send(msg)
	}
}
