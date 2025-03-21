package service

import (
	"context"
	"errors"
	"strings"

	user "github.com/lakhansamani/ecom-grpc-apis/user/v1"
	"github.com/lakhansamani/ecom-grpc-userd/utils"
	"google.golang.org/grpc/metadata"
	"gorm.io/gorm"
)

// Me API to get user details
// Permission: authenticated user
func (s *service) Me(ctx context.Context, req *user.MeRequest) (*user.MeResponse, error) {
	ctx, span := s.trace.Start(ctx, "Me")
	defer span.End()

	// Get the Authorization bearer token from the context
	// Extract the token from the header
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return nil, errors.New("missing authorization token")
	}

	token := authHeader[0]
	// Make sure the token is not empty and is bearer token
	if token == "" {
		meMetrics.WithLabelValues(missingTokenResultLabel).Inc()
		return nil, errors.New("missing token")
	}
	tokenSplit := strings.Split(token, " ")
	if len(tokenSplit) != 2 {
		meMetrics.WithLabelValues(invalidTokenResultLabel).Inc()
		return nil, errors.New("invalid token")
	}
	if strings.ToLower(tokenSplit[0]) != "bearer" {
		meMetrics.WithLabelValues(invalidTokenResultLabel).Inc()
		return nil, errors.New("invalid token")
	}
	token = tokenSplit[1]
	userID, err := utils.VerifyJWT(s.JWTSecret, token)
	if err != nil {
		meMetrics.WithLabelValues(invalidTokenResultLabel).Inc()
		return nil, err
	}
	// Fetch the user from the database
	resUser, err := s.DBProvider.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			meMetrics.WithLabelValues(userNotFoundResultLabel).Inc()
		}
		return nil, err
	}
	meMetrics.WithLabelValues(successResultLabel).Inc()
	// Return the user details
	return &user.MeResponse{
		User: resUser.AsAPIUser(),
	}, nil
}
