package swapi

import (
	"log"
	"strconv"
	"time"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

// MapPersonDTOToDomain converts a SWAPI PersonDTO into domain.Person
func MapPersonDTOToDomain(p PersonDTO) domain.Person {
	created, err := time.Parse(time.RFC3339, p.Created)
	if err != nil {
		log.Printf("warn: failed to parse created date '%s': %v", p.Created, err)
		created = time.Time{}
	}

	mass := parseMass(p.Mass)

	return domain.Person{
		Name:   p.Name,
		Mass:   mass,
		Create: created,
		Films:  p.Films,
	}
}

// MapPlanetDTOToDomain converts a SWAPI PlanetDTO into domain.Planet
func MapPlanetDTOToDomain(pl PlanetDTO) domain.Planet {
	created, err := time.Parse(time.RFC3339, pl.Created)
	if err != nil {
		log.Printf("warn: failed to parse created date '%s': %v", pl.Created, err)
		created = time.Time{}
	}

	return domain.Planet{
		Name:     pl.Name,
		Resident: pl.Residents,
		Created:  created,
		Films:    pl.Films,
	}
}

func parseMass(s string) int {
	if s == "unknown" || s == "" {
		return 0
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("warn: failed to parse mass '%s': %v", s, err)
		return 0
	}
	return v
}
