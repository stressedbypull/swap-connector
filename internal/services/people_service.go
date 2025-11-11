package services

import (
	"context"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/ports"
	"github.com/stressedbypull/swapi-connector/internal/search"
	"github.com/stressedbypull/swapi-connector/internal/sorting"
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

// ListPeople fetches a paginated list of people with search and sorting.
func (s *PeopleService) ListPeople(ctx context.Context, page int, searchTerm, sortBy, sortOrder string) (domain.PaginatedResponse[domain.Person], error) {
	// Fetch from repository
	result, err := s.repo.APIRetrievePeople(ctx, page, searchTerm)
	if err != nil {
		return domain.PaginatedResponse[domain.Person]{}, err
	}

	// Apply search filter
	filtered, err := search.FilterPeopleByName(result.Results, searchTerm)
	if err != nil {
		return domain.PaginatedResponse[domain.Person]{}, err
	}

	// Apply sorting if requested
	if sortBy != "" {
		sorter := sorting.NewPersonSorter(sortBy)
		if sorter != nil {
			ascending := sortOrder == "asc"
			sorter.Sort(filtered, ascending)
		}
	}

	// Update results with filtered and sorted data
	result.Results = filtered
	return result, nil
}

// GetPeopleByID fetches a single person by ID.
func (s *PeopleService) GetPeopleByID(ctx context.Context, id string) (domain.Person, error) {
	return s.repo.APIRetrievePersonByID(ctx, id)
}
