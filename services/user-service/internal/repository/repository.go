package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/goIdioms/gRPC/services/user-service/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (username, email, password, phone, avatar, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	createdAt := time.Now()
	err := r.db.QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password, // Make sure to hash the password before passing it here
		sql.NullString{String: user.Phone, Valid: user.Phone != ""},
		sql.NullString{String: user.Avatar, Valid: user.Avatar != ""},
		createdAt,
	).Scan(
		&user.Id,
		&createdAt,
	)
	user.CreatedAt = createdAt.Unix()

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetUser(id string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) ListUsers() ([]*models.User, error) {
	return nil, nil
}
