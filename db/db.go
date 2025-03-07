package db

import (
	"context"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

// Provider defines the interface for the database provider
type Provider interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
}

// provider implements the Provider interface
type provider struct {
	db *gorm.DB
}

// New creates new database provider
// connects to db and returns the provider
func New(dbURL string) Provider {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		log.Fatalf("Failed to use tracing plugin: %v", err)
	}
	// Auto-migrate User model
	db.AutoMigrate(&User{})

	return &provider{db}
}
