package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func InitDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить .env файл:", err)
	}
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		pgHost := os.Getenv("PGHOST")
		pgPort := os.Getenv("PGPORT")
		pgUser := os.Getenv("PGUSER")
		pgPassword := os.Getenv("PGPASSWORD")
		pgDatabase := os.Getenv("PGDATABASE")

		if pgHost == "" || pgUser == "" || pgDatabase == "" {
			return nil, fmt.Errorf("отсутствуют обязательные переменные окружения (PGHOST, PGUSER, PGDATABASE)")
		}

		if pgPort == "" {
			pgPort = "5432"
		}

		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			pgUser, pgPassword, pgHost, pgPort, pgDatabase)
	}
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
