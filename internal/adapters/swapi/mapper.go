package swapi

import (
	"log"
	"strconv"
	"time"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

func MapPersonToDomain(p PersonDTO) domain.Person {
	person := domain.Person{
		Name:   p.Name,
		Mass:   parseMass(p.Mass),
		Create: parseCreatedDate(p.Create.GoString()),
		Films:  p.Films,
	}
	return person
}

func MapPlanetToDomain(pl PlanetDTO) domain.Planet {
	var planet domain.Planet
	return planet
}

func parseCreatedDate(s string) time.Time {
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		log.Printf("failed to parse date '%s': %v", s, err)
		t = time.Time{}
	}
	return t
}

func parseMass(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("failed to parse mass '%s': %v", s, err)
		value = 9999999
	}
	return value
}
