package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	userpb "github.com/lakhansamani/ecom-grpc-apis/user/v1"

	"github.com/lakhansamani/ecom-grpc-userd/db"
	"github.com/lakhansamani/ecom-grpc-userd/service"
)

func main() {
	// Read .env file as environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using environment variables")
	}

	// DB URL
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is required")
	}
	// JWT Secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}
	// Initialize database
	dbProvider := db.New(dbURL)

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register UserService with gRPC
	userService := service.New(
		service.Config{
			JWTSecret: jwtSecret,
		},
		service.Dependencies{
			DBProvider: dbProvider,
		})
	userpb.RegisterUserServiceServer(server, userService)

	// Start gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("gRPC Server is running on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
