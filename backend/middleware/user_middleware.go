package middleware

import (
	"backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ensures that the request has a valid JWT token
func UserMiddleware(c *gin.Context) {
	// Get the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		c.Abort()
		return
	}

	// Extract the token from the header (format: Bearer <token>)
	tokenString := strings.Split(authHeader, " ")[1]

	// Validate the token and extract the claims
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// Set the user ID from the claims in the context
	c.Set("userID", claims.Username)

	// Proceed with the next handler
	c.Next()
}
