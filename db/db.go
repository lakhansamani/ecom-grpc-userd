package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Provider defines the interface for the database provider
type Provider interface {
	CreateUser(user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id string) (*User, error)
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

	// Auto-migrate User model
	db.AutoMigrate(&User{})

	return &provider{db}
}
