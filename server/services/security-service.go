package services

import (
	"golang.org/x/crypto/bcrypt"
)

type EncryptionService interface {
	EncryptPassword(password string) (string, error)
	ValidatePassword(hashedPassword string, plainTextPassword string) error
}

type BcryptEncryptionService struct{}

func (b *BcryptEncryptionService) EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err	
	}
	return string(bytes), nil
}

func (b *BcryptEncryptionService) ValidatePassword(hashedPassword string, plainTextPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword)); err != nil {
		return err	
	}
	return nil
}