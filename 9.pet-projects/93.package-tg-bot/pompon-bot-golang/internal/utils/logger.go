package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	LevelInfo  = "INFO"
	LevelError = "ERROR"
	LevelDebug = "DEBUG"
)

// LoggerConfig содержит настройки для логгера.
type LoggerConfig struct {
	Prefix string // Префикс для логов
	Level  string // Уровень логирования (INFO, ERROR, DEBUG)
}

// CreateLogger создает логгер с уровнями и настраиваемым префиксом.
func CreateLogger(logFilePath string, config LoggerConfig) *log.Logger {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("[FATAL] Ошибка при создании лог-файла: %v", err)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)
	prefix := config.Prefix
	if prefix == "" {
		prefix = "APP"
	}
	logger := log.New(multiWriter, prefix+": ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}

// LogInfo логирует сообщение с уровнем INFO.
func LogInfo(logger *log.Logger, message string) {
	logger.Printf("[%s] %s", LevelInfo, message)
}

// LogError логирует сообщение с уровнем ERROR.
func LogError(logger *log.Logger, message string) {
	logger.Printf("[%s] %s", LevelError, message)
}

// LogDebug логирует сообщение с уровнем DEBUG.
func LogDebug(logger *log.Logger, message string) {
	logger.Printf("[%s] %s", LevelDebug, message)
}

// LogWithFields логирует сообщение с дополнительными полями.
func LogWithFields(logger *log.Logger, level string, message string, fields map[string]interface{}) {
	fieldsStr := ""
	for key, value := range fields {
		fieldsStr += fmt.Sprintf("%s=%v ", key, value)
	}
	logger.Printf("[%s] %s %s", level, message, fieldsStr)
}
