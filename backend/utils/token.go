package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Define the secret key used for signing the JWT token (should be in your environment variables in production)
var jwtKey = []byte("your_secret_key_here")

// Claims struct defines the structure of the JWT claims
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token for a user
func GenerateToken(username, role string) (string, error) {
	// Set the expiration time of the token (e.g., 1 hour)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims
	claims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "workoutapp",
		},
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatal("Error signing the token:", err)
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(tokenString string) (*Claims, error) {
	// Parse the JWT token and validate it
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract claims from the token
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
