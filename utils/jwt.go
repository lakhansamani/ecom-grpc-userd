package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// Generate JWT token
func GenerateJWT(secret, userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secret))
}

// VerifyJWT verifies the JWT token
func VerifyJWT(secret, tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}
	// Check if token is valid
	if !token.Valid {
		return "", err
	}
	return claims["user_id"].(string), nil
}
