package swapi

import "time"

type PersonDTO struct {
	Name   string    `json:"name"`
	Mass   string    `json:"mass"`
	Create time.Time `json:"created"`
	Films  []string  `json:"films"`
}

type PlanetDTO struct {
	Name     string    `json:"name"`
	Resident []string  `json:"residents"`
	Created  time.Time `json:"created"`
	Films    []string  `json:"films"`
}
