package main

import (
	"log"
	"os"
	"test_case/controllers"
	"test_case/database"
	"test_case/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to the database
	database.Connect()

	// Create a new Gin router
	router := gin.Default()

	// Apply CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Public routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/user", controllers.GetUserData)
		protected.PUT("/user", controllers.UpdateUserAddress)
	}

	// Debugging Routes
	for _, route := range router.Routes() {
		log.Printf("%s %s\n", route.Method, route.Path)
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
