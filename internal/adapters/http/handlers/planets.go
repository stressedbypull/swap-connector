package handlers

import "github.com/stressedbypull/swapi-connector/internal/ports"

type PlanetHandler struct {
	Service ports.PlanetServiceInterface
}

func NewPlanetHandler(s ports.PlanetServiceInterface) *PlanetHandler {
	return &PlanetHandler{
		Service: s,
	}
}
