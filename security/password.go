package security

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes given password with bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks given password with given hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CheckPasswordStructure checks if password has valid format
func CheckPasswordStructure(password string) (err error) {
	if len(password) < 4 || len(password) > 12 {
		err = errors.New("Password length must be between 4 and 12 letters")
	}

	return
}
