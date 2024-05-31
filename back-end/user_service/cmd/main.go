package main

import (
	"log"
	"net"

	"github.com/anh-ngn/group-expense-app/user_service/api/user"
	"github.com/anh-ngn/group-expense-app/user_service/internal/user"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := user.NewUserService()
	user.RegisterUserServiceServer(grpcServer, userService)

	log.Println("Starting gRPC server on :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
