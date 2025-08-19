package main

import (
	"log"
	"os"

	"todo-api/handlers"
	"todo-api/routes"
	"todo-api/services"
)

func main() {
	// Initialize services
	taskService := services.NewTaskService()

	// Initialize handlers
	taskHandler := handlers.NewTaskHandler(taskService)

	// Setup routes
	router := routes.SetupRoutes(taskHandler)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
