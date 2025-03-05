package service

import (
	user "github.com/lakhansamani/ecom-grpc-apis/user/v1"
	"github.com/lakhansamani/ecom-grpc-userd/db"
)

type Config struct {
	// Add configuration here
	JWTSecret string
}

type Dependencies struct {
	// Add dependencies here
	DBProvider db.Provider
}

// Service implements the User service.
type Service interface {
	user.UserServiceServer
}

type service struct {
	Config
	Dependencies
}

// New creates a new User service.
func New(cfg Config, deps Dependencies) Service {
	return &service{
		Config:       cfg,
		Dependencies: deps,
	}
}
