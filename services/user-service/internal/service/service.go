package service

import (
	"context"
	"database/sql"
	"fmt"
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
	fmt.Printf("Received CreateUser request: Name=%s, Email=%s\n", req.Name, req.Email)

	user := &pb.User{
		Id:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now().Unix(),
	}

	return &pb.CreateUserResponse{User: user}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	fmt.Printf("Received GetUser request: ID=%s\n", req.Id)

	user := &pb.User{
		Id:        uuid.New().String(),
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now().Unix(),
	}

	return &pb.GetUserResponse{User: user}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	fmt.Printf("Received ListUsers request: Page=%d, PageSize=%d\n", req.Page, req.PageSize)
	return &pb.ListUsersResponse{}, nil
}
