package db

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	user "github.com/lakhansamani/ecom-grpc-apis/user/v1"
)

// User represents the User model in DB
type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

// AsAPIUser converts the User model to API User
func (u *User) AsAPIUser() *user.User {
	return &user.User{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

// BeforeSave GORM hook to hash password only if it's changed
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Before create
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID
	u.ID = uuid.NewString()
	return
}

// CreateUser creates a new user in the database
func (p *provider) CreateUser(ctx context.Context, u *User) (*User, error) {
	err := p.db.WithContext(ctx).Create(u).Error
	return u, err
}

// GetUserByEmail fetches a user by email from the database
func (p *provider) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	err := p.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return &u, err
}

// GetUserByID fetches a user by ID from the database
func (p *provider) GetUserByID(ctx context.Context, id string) (*User, error) {
	var u User
	err := p.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	return &u, err
}
