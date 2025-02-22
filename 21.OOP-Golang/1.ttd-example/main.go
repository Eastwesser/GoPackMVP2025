package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3" // Импортируем драйвер SQLite
)

// a) Интерфейсы
type MessageSender interface {
	SendMessage(senderID int, receiverID int, text string) error
}

type MessageStorage interface {
	SaveMessage(message Message) error
	GetMessages(userID int) ([]Message, error)
}

// b) Структуры
type Message struct {
	ID         int
	SenderID   int
	ReceiverID int
	Text       string
	Timestamp  time.Time
}

// c) Реализация
type SQLiteMessageStorage struct {
	db *sql.DB
}

func (s SQLiteMessageStorage) GetMessages(userID int) ([]Message, error) {
	query := "SELECT id, sender_id, receiver_id, text, timestamp FROM messages WHERE receiver_id = ?"
	rows, err := s.db.Query(
		query,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(
			&msg.ID,
			&msg.SenderID,
			&msg.ReceiverID,
			&msg.Text,
			&msg.Timestamp,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (s SQLiteMessageStorage) SaveMessage(message Message) error {
	_, err := s.db.Exec("INSERT INTO messages (sender_id, receiver_id, text, timestamp) VALUES (?, ?, ?, ?)",
		message.SenderID, message.ReceiverID, message.Text, message.Timestamp)
	return err
}

type MessageService struct {
	storage MessageStorage
}

func (s MessageService) SendMessage(senderID int, receiverID int, text string) error {
	message := Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Text:       text,
		Timestamp:  time.Now(),
	}
	return s.storage.SaveMessage(message)
}

// Создание таблицы messages
func createMessagesTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sender_id INTEGER NOT NULL,
		receiver_id INTEGER NOT NULL,
		text TEXT NOT NULL,
		timestamp DATETIME NOT NULL
	);`
	_, err := db.Exec(query)
	return err
}

// d) Использование
func main() {
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		fmt.Println("Ошибка при открытии базы данных:", err)
		return
	}
	defer db.Close()

	// Создаем таблицу messages
	if err := createMessagesTable(db); err != nil {
		fmt.Println("Ошибка при создании таблицы:", err)
		return
	}

	storage := SQLiteMessageStorage{db: db}
	service := MessageService{storage: storage}

	// Отправляем сообщение
	err = service.SendMessage(1, 2, "Привет!")
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения:", err)
	}

	// Получаем сообщения для пользователя с ID 2
	messages, err := storage.GetMessages(2)
	if err != nil {
		fmt.Println("Ошибка при получении сообщений:", err)
		return
	}

	// Выводим сообщения
	fmt.Println("Сообщения для пользователя 2:")
	for _, msg := range messages {
		fmt.Printf("От: %d, Текст: %s, Время: %s\n",
			msg.SenderID,
			msg.Text,
			msg.Timestamp,
		)
	}
}
