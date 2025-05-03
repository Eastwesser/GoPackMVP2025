package handlers

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"pompon-bot-golang/internal/keyboards"
	"pompon-bot-golang/internal/models"
	"pompon-bot-golang/internal/services"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type OrderService struct {
	db *gorm.DB
}

// userState stores user state during the order process
var userState = make(map[int64]string)

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

// GetCategories возвращает список категорий товаров
func (s *OrderService) GetCategories(ctx context.Context) ([]string, error) {
	// Запрашиваем список категорий из базы данных
	var categories []models.Category
	if err := s.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}

	// Преобразуем категории в строковый срез
	var categoryNames []string
	for _, category := range categories {
		categoryNames = append(categoryNames, category.Name)
	}

	return categoryNames, nil
}

func HandleOrder(orderService *services.OrderService, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageText := update.Message.Text

	log.Printf("Processing order for chat_id: %d", chatID)

	if messageText == "/order" {
		categories, err := orderService.GetCategories(context.Background())
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, "Error loading categories.")
			bot.Send(msg)
			return
		}

		msg := tgbotapi.NewMessage(chatID, "Please choose a category:")
		msg.ReplyMarkup = keyboards.CatalogKeyboard(categories) // Send categories
		bot.Send(msg)
		userState[chatID] = "waiting_for_category"
		return
	}

	switch userState[chatID] {
	case "waiting_for_category":
		msg := tgbotapi.NewMessage(chatID, "Please enter the quantity of items:")
		bot.Send(msg)
		userState[chatID] = "waiting_for_quantity"
	case "waiting_for_quantity":
		quantity, err := strconv.Atoi(messageText)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, "Please enter a valid number.")
			bot.Send(msg)
			return
		}
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Your order for %d items has been accepted! 🎉", quantity))
		bot.Send(msg)
		delete(userState, chatID)
	}
}
