package controllers

import (
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUeser(c *gin.Context) {
	var Input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	hashedPassword, err := utils.HashPassword(Input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	var role models.Role
	if err := models.DB.Where("name = ?", "User").First(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User role not found"})
		return
	}

	user := models.User{
		Username: Input.Username,
		Email:    Input.Email,
		Password: hashedPassword,
		RoleID:   role.ID,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func loginUser(c *gin.Context) {
	var Input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
}
