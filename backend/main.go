package main

import (
	"backend/models"
	"backend/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	models.InitDB()

	// Set up the Gin router
	r := gin.Default()

	// Define a root route that returns a basic message
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go server.",
		})
	})

	// Register user-related routes
	routes.SetupUserRoutes(r)

	// Log the server URL
	log.Println("Server is running at http://localhost:8080")

	// Run the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
