package services

import (
	"context"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/ports"
)

type PeopleService struct {
	repo ports.PeopleRepository
}

func NewPeopleService(r ports.PeopleRepository) *PeopleService {
	return &PeopleService{repo: r}
}

func (p *PeopleService) ListPeople(ctx context.Context, page, pageSize int, search, sortBy, sortOrder string) (domain.PaginatedResponse[domain.Person], error) {
	// TODO: Implement sorting and filtering
	return p.repo.APIRetrievePeople(ctx, page, search)
}

func (p *PeopleService) GetPeopleByID(ctx context.Context, id string) (domain.Person, error) {
	return p.repo.APIRetrievePersonByID(ctx, id)
}
