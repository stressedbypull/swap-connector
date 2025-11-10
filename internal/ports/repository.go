package ports

import (
	"context"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

// PeopleRepository is a port for fetching people
type PeopleRepository interface {
	FetchPeople(ctx context.Context, page, pageSize int, search string) (domain.PaginatedResponse[domain.Person], error)
	//FetchPersonByID(ctx context.Context, id string) (domain.Person, error)
}

// PlanetsRepository is a port for fetching planets
type PlanetsRepository interface {
	FetchPlanets(ctx context.Context) ([]domain.Planet, error)
	FetchPlanetByID(ctx context.Context, id string) (domain.Planet, error)
}
