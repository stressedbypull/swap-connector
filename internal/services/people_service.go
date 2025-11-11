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

func (p *PeopleService) ListPeople(ctx context.Context, page, pageSize int, search string) (domain.PaginatedResponse[domain.Person], error) {
	return domain.PaginatedResponse[domain.Person]{}, nil
}
