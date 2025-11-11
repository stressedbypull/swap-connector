package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/handlers"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/middleware"
	"github.com/stressedbypull/swapi-connector/internal/adapters/swapi"
	"github.com/stressedbypull/swapi-connector/internal/services"
)

const (
	swapiBaseURL = "https://swapi.dev/api"
	serverPort   = ":6969"
)

func main() {
	// Dependency Injection: Infrastructure -> Adapter -> Service -> Handler

	// 1. Infrastructure layer: HTTP client
	httpClient := &http.Client{}

	// 2. Adapter layer: SWAPI client implements repository interfaces
	swapiClient := swapi.NewClient(swapiBaseURL, httpClient)

	// 3. Service layer: Business logic
	peopleService := services.NewPeopleService(swapiClient)

	// 4. Presentation layer: HTTP handlers
	peopleHandler := handlers.NewPeopleHandler(peopleService)

	// Setup router
	router := gin.Default()

	// Global middleware
	router.Use(middleware.PaginationMiddleware())
	router.Use(middleware.QueryMiddleware())

	// Health check
	router.GET("/ping", healthCheck)

	// API route groups
	api := router.Group("/api")
	{
		// People endpoints
		people := api.Group("/people")
		{
			people.GET("", peopleHandler.ListPeople)
		}

		// Planets endpoints (TODO: implement)
		// planets := api.Group("/planets")
		// {
		// 	planets.GET("", planetsHandler.ListPlanets)
		// }
	}

	// Start server
	log.Printf("Starting server on port %s", serverPort)
	if err := router.Run(serverPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// healthCheck handles health check requests.
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "pong",
	})
}
