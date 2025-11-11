package services

import (
	"context"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/ports"
)

// PeopleService handles business logic for people operations.
type PeopleService struct {
	repo ports.PeopleRepository
}

// NewPeopleService creates a new people service with dependency injection.
func NewPeopleService(repo ports.PeopleRepository) *PeopleService {
	return &PeopleService{
		repo: repo,
	}
}

// ListPeople fetches a paginated list of people.
// SWAPI controls pagination (10 items per page).
// TODO: Implement sorting and search filtering.
func (s *PeopleService) ListPeople(ctx context.Context, page int, search, sortBy, sortOrder string) (domain.PaginatedResponse[domain.Person], error) {
	// For now, just pass through to repository
	// In the future, apply sorting and filtering here
	return s.repo.APIRetrievePeople(ctx, page, search)
}

// GetPeopleByID fetches a single person by ID.
func (s *PeopleService) GetPeopleByID(ctx context.Context, id string) (domain.Person, error) {
	return s.repo.APIRetrievePersonByID(ctx, id)
}
