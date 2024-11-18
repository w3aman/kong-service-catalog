package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/w3aman/kong-service-catalog/api"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Middleware for logging
	router.Use(middleware.LoggingMiddleware)

	// Register API routes
	api.RegisterRoutes(router)

	// Start the server
	log.Println("Starting server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
