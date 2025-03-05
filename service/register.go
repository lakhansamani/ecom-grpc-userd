package service

import (
	"context"
	"errors"
	"strings"

	user "github.com/lakhansamani/ecom-grpc-apis/user/v1"

	"github.com/lakhansamani/ecom-grpc-userd/db"
)

// Register API to register a new user
// Permission: none
func (s *service) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	name := req.GetName()
	email := req.GetEmail()
	password := req.GetPassword()

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name is required")
	}

	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email is required")
	}

	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password is required")
	}

	resUser, err := s.DBProvider.CreateUser(&db.User{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{
		UserId: resUser.ID,
	}, nil
}
