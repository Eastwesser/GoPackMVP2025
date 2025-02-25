package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Настройки приложения

type Config struct {
	TelegramAPIKey string
	DatabaseURL    string
	DBUser         string
	DBPassword     string
	DBName         string
	DBHost         string
	DBPort         string
}

func LoadConfig() *Config {
	// Загрузка из .env
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := &Config{}

	cfg.DatabaseURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	cfg.TelegramAPIKey = os.Getenv("TELEGRAM_API_KEY")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")

	return cfg
}
