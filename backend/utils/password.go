package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain-text password using bcrypt
func HashPassword(password string) (string, error) {
	// Generate a hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a hashed password with a plain-text password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
