package mocks

import (
	"context"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockSwapiRepository struct {
	mock.Mock
}

func NewMockSwapiRepository() *MockSwapiRepository {
	var mock MockSwapiRepository
	return &mock
}

func (m *MockSwapiRepository) APIRetrievePeople(ctx context.Context, page int, search string) (domain.PaginatedResponse[domain.Person], error) {
	args := m.Called(ctx, page, search)
	var paginatedPeople domain.PaginatedResponse[domain.Person]
	if args.Get(0) != nil {
		paginatedPeople = args.Get(0).(domain.PaginatedResponse[domain.Person])
	}
	return paginatedPeople, nil
}
