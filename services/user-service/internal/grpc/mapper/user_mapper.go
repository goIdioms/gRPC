package mapper

import (
	pb "github.com/goIdioms/gRPC/api/proto/user/v1"
	"github.com/goIdioms/gRPC/services/user-service/internal/models"
)

func ProtoToDomainUser(protoUser *pb.User) *models.User {
	if protoUser == nil {
		return nil
	}

	return &models.User{
		Id:        protoUser.Id,
		Username:  protoUser.Username,
		Email:     protoUser.Email,
		Password:  protoUser.Password,
		Phone:     protoUser.Phone,
		CreatedAt: protoUser.CreatedAt,
	}
}

func DomainToProtoUser(domainUser *models.User) *pb.User {
	if domainUser == nil {
		return nil
	}

	return &pb.User{
		Id:        domainUser.Id,
		Username:  domainUser.Username,
		Email:     domainUser.Email,
		CreatedAt: domainUser.CreatedAt,
	}
}
