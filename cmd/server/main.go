package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/handlers"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/middleware"
	"github.com/stressedbypull/swapi-connector/internal/adapters/swapi"
	"github.com/stressedbypull/swapi-connector/internal/config"
	"github.com/stressedbypull/swapi-connector/internal/services"

	_ "github.com/stressedbypull/swapi-connector/docs" // Import generated docs
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           SWAPI Connector API
// @version         1.0
// @description     A Star Wars API connector with pagination, search, and sorting capabilities
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@swapi-connector.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:6969
// @BasePath  /api

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	// Load configuration from environment variables
	cfg := config.Load()

	// Dependency Injection: Infrastructure -> Adapter -> Service -> Handler

	// 1. Infrastructure layer: HTTP client
	httpClient := &http.Client{}

	// 2. Adapter layer: SWAPI client implements repository interfaces
	swapiClient := swapi.NewClient(cfg.SWAPI.BaseURL, httpClient)

	// 3. Service layer: Business logic
	peopleService := services.NewPeopleService(swapiClient)

	// 4. Presentation layer: HTTP handlers
	peopleHandler := handlers.NewPeopleHandler(peopleService)

	// Setup router
	router := gin.Default()

	// Global middleware
	router.Use(middleware.CORS(cfg.CORS.AllowedOrigins))
	router.Use(middleware.PaginationMiddleware())
	router.Use(middleware.QueryMiddleware())

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	log.Printf("Starting server on port %s", cfg.Server.Port)
	if err := router.Run(cfg.Server.Port); err != nil {
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
