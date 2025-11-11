package handlers

import "github.com/stressedbypull/swapi-connector/internal/ports"

type PeopleHandler struct {
	Service ports.PeopleServiceInterface
}

func NewPeopleHandler(s ports.PeopleServiceInterface) *PeopleHandler {
	return &PeopleHandler{
		Service: s,
	}
}
