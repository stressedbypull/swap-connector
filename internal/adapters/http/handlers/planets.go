package handlers

import "github.com/stressedbypull/swapi-connector/internal/ports"

// PlanetHandler handles HTTP requests for planet resources.
type PlanetHandler struct {
	service ports.PlanetServiceInterface // Use lowercase for consistency with PeopleHandler
}

// NewPlanetHandler creates a new planet handler with dependency injection.
func NewPlanetHandler(service ports.PlanetServiceInterface) *PlanetHandler {
	return &PlanetHandler{
		service: service,
	}
}
