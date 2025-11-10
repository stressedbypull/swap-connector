package services

import (
	"context"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/ports"
)

type PlanetService struct {
	repo ports.PlanetsRepository
}

func NewPlanetService(r ports.PlanetsRepository) *PlanetService {
	return &PlanetService{repo: r}
}

func (s *PlanetService) GetAll(ctx context.Context) ([]domain.Planet, error) {
	return s.repo.FetchPlanets(ctx)
}

func (s *PlanetService) GetByID(ctx context.Context, id string) (domain.Planet, error) {
	return s.repo.FetchPlanetByID(ctx, id)
}
