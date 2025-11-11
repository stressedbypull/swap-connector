package ports

import (
	"context"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

// PeopleService - Interface for business logic
type PeopleServiceInterface interface {
	ListPeople(ctx context.Context, page, pageSize int, search, sortBy, sortOrder string) (domain.PaginatedResponse[domain.Person], error)
	GetPeopleByID(ctx context.Context, id string) (domain.Person, error)
}

// PlanetService - Interface for business logic
type PlanetServiceInterface interface {
	ListPlanets(ctx context.Context, page, pageSize int, search, sortBy, sortOrder string) (domain.PaginatedResponse[domain.Planet], error)
	GetPlanetByID(ctx context.Context, id string) (domain.Planet, error)
}
