package services

import (
	"github.com/stressedbypull/swapi-connector/internal/ports"
)

type PeopleService struct {
	repo ports.PeopleRepository
}

func NewPeopleService(r ports.PeopleRepository) *PeopleService {
	return &PeopleService{repo: r}
}
