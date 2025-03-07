package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	userpb "github.com/lakhansamani/ecom-grpc-apis/user/v1"

	"github.com/lakhansamani/ecom-grpc-userd/db"
	"github.com/lakhansamani/ecom-grpc-userd/service"
)

const (
	serviceName = "userd"
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

	ctx := context.Background()
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint("localhost:4317"))
	if err != nil {
		log.Fatalf("failed to create OTLP trace exporter: %v", err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	defer func() {
		tracerProvider.Shutdown(context.Background())
	}()

	otel.SetTracerProvider(tracerProvider)

	// Create OpenTelemetry gRPC stats handler
	serverHandler := otelgrpc.NewServerHandler(
		otelgrpc.WithTracerProvider(tracerProvider),
	)

	// Create a new gRPC server
	server := grpc.NewServer(
		grpc.StatsHandler(serverHandler),
	)
	// Register UserService with gRPC
	userService := service.New(
		service.Config{
			JWTSecret: jwtSecret,
		},
		service.Dependencies{
			DBProvider:    dbProvider,
			TraceProvider: tracerProvider,
		})
	userpb.RegisterUserServiceServer(server, userService)

	// Start gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	reflection.Register(server)
	log.Println("gRPC Server is running on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
