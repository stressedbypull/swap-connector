package services

import (
	"context"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/ports"
	"github.com/stressedbypull/swapi-connector/internal/search"
	"github.com/stressedbypull/swapi-connector/internal/sorting"
)

type PlanetService struct {
	repo ports.PlanetsRepository
}

func NewPlanetService(r ports.PlanetsRepository) *PlanetService {
	return &PlanetService{repo: r}
}

// ListPlanets fetches a paginated list of planets with search and sorting.
func (s *PlanetService) ListPlanets(ctx context.Context, page int, searchTerm, sortBy, sortOrder string) (domain.PaginatedResponse[domain.Planet], error) {
	// Fetch from repository
	result, err := s.repo.FetchPlanets(ctx, page, searchTerm)
	if err != nil {
		return domain.PaginatedResponse[domain.Planet]{}, err
	}

	// Apply search filter
	result.Results = search.FilterPlanetsByName(result.Results, searchTerm)

	// Apply sorting if requested
	if sortBy != "" {
		sorter := sorting.NewPlanetSorter(sortBy)
		if sorter != nil {
			ascending := sortOrder == "asc"
			sorter.Sort(result.Results, ascending)
		}
	}

	return result, nil
}

// GetPlanetByID fetches a single planet by ID.
func (s *PlanetService) GetPlanetByID(ctx context.Context, id string) (domain.Planet, error) {
	return s.repo.FetchPlanetByID(ctx, id)
}
