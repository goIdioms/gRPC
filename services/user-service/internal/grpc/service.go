package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	pb "github.com/goIdioms/gRPC/api/proto/user/v1"
	"github.com/goIdioms/gRPC/services/user-service/internal/grpc/mapper"
	"github.com/goIdioms/gRPC/services/user-service/internal/repository"
	"github.com/google/uuid"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	domainUser := mapper.ProtoToDomainUser(&pb.User{
		Id:        uuid.New().String(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		Phone:     req.Phone,
		CreatedAt: time.Now().Unix(),
	})
	log.Println(domainUser)

	user, err := s.repo.CreateUser(domainUser)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	log.Printf("Successfully created user with ID: %s", user.Id)
	return &pb.CreateUserResponse{
		User: mapper.DomainToProtoUser(user),
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.repo.GetUser(req.Id)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	return &pb.GetUserResponse{
		User: mapper.DomainToProtoUser(user),
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {

	users, err := s.repo.ListUsers()
	if err != nil {
		return nil, errors.New("failed to list users")
	}

	var protoUsers []*pb.User
	for _, user := range users {
		protoUsers = append(protoUsers, mapper.DomainToProtoUser(user))
	}

	return &pb.ListUsersResponse{
		Users: protoUsers,
	}, nil
}
