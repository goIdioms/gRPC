package service

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/goIdioms/gRPC/api/proto/user/v1"
	"github.com/google/uuid"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	user := &pb.User{
		Id:        uuid.New().String(),
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: time.Now().Unix(),
	}

	return &pb.CreateUserResponse{User: user}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	user := &pb.User{
		Id:        uuid.New().String(),
		Username:  "John Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now().Unix(),
	}

	return &pb.GetUserResponse{User: user}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return &pb.ListUsersResponse{}, nil
}
