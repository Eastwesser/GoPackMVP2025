package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectDB инициализирует подключение к PostgreSQL
func ConnectDB(cfg *Config) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := dbpool.Ping(ctx); err != nil {
		log.Fatalf("Ошибка пинга базы данных: %v", err)
	}

	log.Println("Успешное подключение к базе данных!")
	return dbpool
}
