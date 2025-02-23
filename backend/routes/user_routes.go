package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

// UserRoutes handles user-related routes
func UserRoutes(r *gin.Engine) {
	// Public routes (no authentication required)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes (authentication required)
	r.GET("/me", middleware.UserMiddleware, controllers.GetCurrentUser)
	r.PUT("/me", middleware.UserMiddleware, controllers.UpdateUser)
}
