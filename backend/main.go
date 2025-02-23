package main

import (
	"backend/config"
	"backend/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	if err := config.InitDB(); err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	// Set up the Gin router
	r := gin.Default()

	// Define a root route that returns a basic message
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go server.",
		})
	})

	// Register user-related routes
	routes.UserRoutes(r)

	// You can add other route groups like workout routes, etc., in the future
	// routes.SetupWorkoutRoutes(r)

	// Log the server URL
	log.Printf("Server is running at http://localhost:8080")

	// Run the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
