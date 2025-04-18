package services

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type TokenService interface {
	GenerateToken(subject string) (string, error)
}

type JwtService struct{}

func (j *JwtService) GenerateToken(subject string) (string, error) {
	secret := []byte(os.Getenv("jwt-secret-key"))
	claims := jwt.StandardClaims{
		Subject: subject,
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		IssuedAt: time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}