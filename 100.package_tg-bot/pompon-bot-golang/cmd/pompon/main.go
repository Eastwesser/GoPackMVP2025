package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"pompon-bot-golang/internal/config"
	"pompon-bot-golang/internal/handlers"
	"pompon-bot-golang/internal/models"
	"pompon-bot-golang/internal/services"
	"pompon-bot-golang/internal/utils"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var logger *log.Logger

func init() {
	// Создаем логгер один раз при старте приложения
	logger = utils.CreateLogger("bot.log", utils.LoggerConfig{
		Prefix: "POMPON-BOT",
		Level:  utils.LevelInfo,
	})
}

func main() {
	utils.LogInfo(logger, "Бот запускается...")

	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключение к базе данных
	db := config.ConnectDB(cfg)
	defer db.Close()

	// Инициализация Telegram Bot API
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPIKey)
	if err != nil {
		utils.LogError(logger, fmt.Sprintf("Failed to initialize bot: %v", err))
		os.Exit(1)
	}

	bot.Debug = true
	utils.LogInfo(logger, fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	// Инициализация сервисов
	catalogService := services.NewCatalogService(db)
	orderService := services.NewOrderService(db) // Убедитесь, что в NewOrderService используется правильный тип для базы данных
	notifyService := services.NewNotifyService(db, bot)
	subscribeService := services.NewSubscribeService(db)

	// Канал для обработки системных сигналов
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Канал для завершения горутины обработки обновлений
	done := make(chan struct{})

	// Обработка обновлений Telegram в горутине
	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message != nil {
				if update.Message.IsCommand() {
					switch update.Message.Command() {
					case "start":
						handlers.HandleStart(update, bot)
					case "about":
						handlers.HandleAbout(update, bot)
					case "catalog":
						handlers.HandleCatalog(catalogService, bot, update.Message.Chat.ID)
					case "order":
						handlers.HandleOrder(orderService, bot, update)
					case "subscribe":
						handlers.HandleSubscribe(subscribeService, bot, update.Message.Chat.ID)
					default:
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Используйте /start для начала.")
						bot.Send(msg)
					}
				} else {
					// Обработка текстовых сообщений
					switch update.Message.Text {
					case "🔹 О нас", "🔹 About Us":
						handlers.HandleAbout(update, bot)
					case "📦 Каталог", "📦 Catalog":
						handlers.HandleCatalog(catalogService, bot, update.Message.Chat.ID)
					case "🛒 Заказать", "🛒 Order":
						handlers.HandleOrder(orderService, bot, update)
					case "🔔 Подписка", "🔔 Subscription":
						handlers.HandleSubscribe(subscribeService, bot, update.Message.Chat.ID)
					default:
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Используйте /start для начала.")
						bot.Send(msg)
					}
				}

				// Пример использования сервисов внутри горутины
				if update.Message.Text == "🛒 Заказать" {
					order := models.Order{
						UserID:   update.Message.Chat.ID, // Используем chat_id пользователя
						Product:  models.Product{ID: 1},
						Quantity: 2,
						Status:   "pending",
					}

					if err := orderService.CreateOrder(context.Background(), order.UserID, order); err != nil {
						utils.LogError(logger, fmt.Sprintf("Failed to create order: %v", err))
					} else {
						// Отправляем уведомление подписчикам о новом заказе
						message := fmt.Sprintf("Новый заказ создан: %s, количество: %d", order.Product.Name, order.Quantity)
						if err := notifyService.NotifyAllSubscribers(context.Background(), message); err != nil {
							utils.LogError(logger, fmt.Sprintf("Failed to notify subscribers: %v", err))
						}
					}
				}

			} else if update.CallbackQuery != nil {
				// Обработка callback-запросов (например, нажатие на кнопки)
				handlers.HandleCallbackQuery(db, bot, update.CallbackQuery)
			}
		}
		done <- struct{}{}
	}()

	// Ожидание сигнала завершения
	<-stop
	utils.LogInfo(logger, "Получен сигнал завершения работы, закрытие бота...")

	// Ожидаем завершения горутины обработки обновлений
	<-done
	utils.LogInfo(logger, "Бот завершает работу.")
}
