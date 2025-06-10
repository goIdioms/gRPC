package service

import (
	"context"
	"time"

	pb "github.com/goIdioms/gRPC/api/proto/user/v1"
	"github.com/goIdioms/gRPC/services/user-service/internal/grpc/mapper"
	"github.com/goIdioms/gRPC/services/user-service/internal/models"
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
		Username: req.Username,
		Email:    req.Email,
	})

	if err := s.repo.CreateUser(domainUser); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{User: &pb.User{
		Id:        domainUser.Id,
		Username:  domainUser.Username,
		Email:     domainUser.Email,
		CreatedAt: domainUser.CreatedAt,
	}}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	user := &models.User{
		Id:        uuid.New().String(),
		Username:  "John Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now().Unix(),
	}

	return &pb.GetUserResponse{User: &pb.User{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return &pb.ListUsersResponse{}, nil
}
