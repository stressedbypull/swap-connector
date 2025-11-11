package mocks

import (
	"context"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockSwapiRepository is a mock implementation of ports.PeopleRepository
type MockSwapiRepository struct {
	mock.Mock
}

// NewMockSwapiRepository creates a new mock repository
func NewMockSwapiRepository() *MockSwapiRepository {
	return &MockSwapiRepository{}
}

// APIRetrievePeople mocks fetching people with pagination
func (m *MockSwapiRepository) APIRetrievePeople(ctx context.Context, page int, search string) (domain.PaginatedResponse[domain.Person], error) {
	args := m.Called(ctx, page, search)
	return args.Get(0).(domain.PaginatedResponse[domain.Person]), args.Error(1)
}

// APIRetrievePersonByID mocks fetching a single person by ID
func (m *MockSwapiRepository) APIRetrievePersonByID(ctx context.Context, id string) (domain.Person, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Person), args.Error(1)
}
