package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3" // Импортируем драйвер SQLite
)

// Интерфейсы
type MessageSender interface {
	SendMessage(senderID int, receiverID int, text string) error
}

type MessageStorage interface {
	SaveMessage(message Message) error
	GetMessages(userID int) ([]Message, error)
}

// Структуры
type Message struct {
	ID         int
	SenderID   int
	ReceiverID int
	Text       string
	Timestamp  time.Time
}

// Реализация хранилища
type SQLiteMessageStorage struct {
	db *sql.DB
}

func (s SQLiteMessageStorage) SaveMessage(message Message) error {
	_, err := s.db.Exec("INSERT INTO messages (sender_id, receiver_id, text, timestamp) VALUES (?, ?, ?, ?)",
		message.SenderID, message.ReceiverID, message.Text, message.Timestamp)
	return err
}

func (s SQLiteMessageStorage) GetMessages(userID int) ([]Message, error) {
	rows, err := s.db.Query("SELECT id, sender_id, receiver_id, text, timestamp FROM messages WHERE receiver_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Text, &msg.Timestamp)
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

// Реализация сервиса
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

// Создание таблицы
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

/*
Зафиксируй acceptance criteria (критерии приемки). Например:

Пользователь может отправить сообщение.

Сообщение сохраняется в базе данных.

Пользователь может получить свои сообщения.
*/

// Тест
func TestSendMessage(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Ошибка при открытии базы данных: %v", err)
	}
	defer db.Close()

	if err := createMessagesTable(db); err != nil {
		t.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	storage := SQLiteMessageStorage{db: db}
	service := MessageService{storage: storage}

	// Отправка сообщения
	err = service.SendMessage(1, 2, "Привет!")
	if err != nil {
		t.Errorf("Ошибка при отправке сообщения: %v", err)
	}

	// Получение сообщений
	messages, err := storage.GetMessages(2)
	if err != nil {
		t.Errorf("Ошибка при получении сообщений: %v", err)
	}

	// Проверки
	if len(messages) != 1 {
		t.Errorf("Ожидалось 1 сообщение, получено %d", len(messages))
	}

	if messages[0].Text != "Привет!" {
		t.Errorf("Ожидалось сообщение 'Привет!', получено '%s'", messages[0].Text)
	}

	if messages[0].SenderID != 1 {
		t.Errorf("Ожидался отправитель с ID 1, получено %d", messages[0].SenderID)
	}

	if messages[0].ReceiverID != 2 {
		t.Errorf("Ожидался получатель с ID 2, получено %d", messages[0].ReceiverID)
	}
}

// main (для демонстрации)
func main() {
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		fmt.Println("Ошибка при открытии базы данных:", err)
		return
	}
	defer db.Close()

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

	// Получаем сообщения
	messages, err := storage.GetMessages(2)
	if err != nil {
		fmt.Println("Ошибка при получении сообщений:", err)
		return
	}

	// Выводим сообщения
	fmt.Println("Сообщения для пользователя 2:")
	for _, msg := range messages {
		fmt.Printf("От: %d, Текст: %s, Время: %s\n", msg.SenderID, msg.Text, msg.Timestamp)
	}
}
