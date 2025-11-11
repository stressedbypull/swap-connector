package ports

import (
	"context"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

// PeopleRepository is a port for fetching people
type PeopleRepository interface {
	APIRetrievePeople(ctx context.Context, page int, search string) (domain.PaginatedResponse[domain.Person], error)
	APIRetrievePersonByID(ctx context.Context, id string) (domain.Person, error)
}

// PlanetsRepository is a port for fetching planets
type PlanetsRepository interface {
	FetchPlanets(ctx context.Context, page int, search string) (domain.PaginatedResponse[domain.Planet], error)
	FetchPlanetByID(ctx context.Context, id string) (domain.Planet, error)
}
