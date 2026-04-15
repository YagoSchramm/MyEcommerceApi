package service

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateID() *uuid.UUID {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil
	}
	return &id
}
func GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	passwordHash := string(hashedPassword)
	return passwordHash, err
}
func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
