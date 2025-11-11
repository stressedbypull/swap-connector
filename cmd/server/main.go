package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/handlers"
	"github.com/stressedbypull/swapi-connector/internal/adapters/swapi"
	"github.com/stressedbypull/swapi-connector/internal/services"
)

func main() {
	// Dependency Injection Setup
	// 1. Infrastructure: HTTP client (can be mocked for testing)
	httpClient := &http.Client{}

	// 2. Adapter: SWAPI Client implements PeopleRepository & PlanetsRepository
	swapiClient := swapi.NewClient("https://swapi.dev/api", httpClient)

	// 3. Service: Business logic layer
	peopleService := services.NewPeopleService(swapiClient)

	// 4. Presentation: HTTP handlers
	peopleHandler := handlers.NewPeopleHandler(peopleService)

	// Router setup
	router := gin.Default()

	// Health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "pong",
		})
	})

	// People routes
	router.GET("/people", peopleHandler.ListPeople)

	// Start server on port 6969
	router.Run(":6969")
}
