package services

import (
	"golang.org/x/crypto/bcrypt"
)

type EncryptionService interface {
	EncryptPassword(password string) (string, error)
}

type BcryptEncryptionService struct{}

func (b *BcryptEncryptionService) EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err	
	}
	return string(bytes), nil
}