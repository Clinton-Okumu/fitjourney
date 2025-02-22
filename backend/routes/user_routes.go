package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes initializes all user-related routes
func SetupUserRoutes(router *gin.Engine) {
	// Public routes (accessible without authentication)
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Protected routes (require authentication)
	protected := router.Group("/user")
	protected.Use(middleware.AuthMiddleware()) // Apply JWT middleware
	protected.GET("/profile", controllers.GetUserProfile)
}
