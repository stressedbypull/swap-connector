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

func (s *PlanetService) ListPlanets(ctx context.Context, page, pageSize int, search string) (domain.PaginatedResponse[domain.Planet], error) {
	return s.repo.FetchPlanets(ctx, page, search)
}

func (s *PlanetService) GetByID(ctx context.Context, id string) (domain.Planet, error) {
	return s.repo.FetchPlanetByID(ctx, id)
}
