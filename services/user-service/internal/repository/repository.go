package repository

import (
	"database/sql"

	"github.com/goIdioms/gRPC/services/user-service/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return nil
}
