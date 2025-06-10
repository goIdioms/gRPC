package repository

import (
	"database/sql"
	"errors"

	"github.com/goIdioms/gRPC/services/user-service/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {

	err := r.db.QueryRow("INSERT INTO users (id, username, email, password_hash, phone, avatar, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.Id, user.Username, user.Email, user.Password, user.Phone, user.Avatar, user.CreatedAt)
	if err != nil {
		return nil, errors.New("failed to create user")
	}
	return user, nil
}

func (r *UserRepository) GetUser(id string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) ListUsers() ([]*models.User, error) {
	return nil, nil
}
