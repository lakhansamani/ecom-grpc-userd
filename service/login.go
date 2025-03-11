package service

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	user "github.com/lakhansamani/ecom-grpc-apis/user/v1"

	"github.com/lakhansamani/ecom-grpc-userd/utils"
)

// Login API to login a user
// Permission: none
func (s *service) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	ctx, span := s.trace.Start(ctx, "Login")
	defer span.End()

	email := req.GetEmail()
	password := req.GetPassword()

	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email is required")
	}

	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password is required")
	}

	// Get user by email
	resUser, err := s.DBProvider.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			loginMetrics.WithLabelValues(userNotFoundResultLabel).Inc()
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Match password
	if err := bcrypt.CompareHashAndPassword([]byte(resUser.Password), []byte(password)); err != nil {
		loginMetrics.WithLabelValues(invalidPasswordResultLabel).Inc()
		return nil, errors.New("invalid password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(s.JWTSecret, resUser.ID)
	if err != nil {
		return nil, err
	}
	loginMetrics.WithLabelValues(successResultLabel).Inc()
	return &user.LoginResponse{
		Token: token,
	}, nil
}
