package main

import (
	"log"
	"net"

	pb "github.com/goIdioms/gRPC/api/proto/user/v1"
	"github.com/goIdioms/gRPC/pkg/database"
	grpcService "github.com/goIdioms/gRPC/services/user-service/internal/grpc"
	"github.com/goIdioms/gRPC/services/user-service/internal/repository"
	"google.golang.org/grpc"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	userService := grpcService.NewUserService(repo)

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, userService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
