package main

import (
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/routes"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.Connect()

	// AutoMigrate the models
	err := database.DB.AutoMigrate(&models.User{}, &models.Course{}) // Add Course model migration
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	// Set up router
	router := gin.Default()

	// Register routes
	routes.UserRoutes(router)
	// routes.CourseRoutes(router) // Register course routes

	// Run the server
	router.Run(":8080")
}
