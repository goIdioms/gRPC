package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/goIdioms/gRPC/pkg/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB() (*sql.DB, error) {
	config, err := config.LoadConfig("/Users/az/Desktop/gRPC/config")
	if err != nil {
		return nil, fmt.Errorf("ошибка загрузки конфигурации: %w", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия соединения с БД: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}

	log.Println("Успешное подключение к PostgreSQL!")
	return db, nil
}
