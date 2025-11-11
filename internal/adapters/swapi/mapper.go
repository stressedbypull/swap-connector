package swapi

import (
	"log"
	"time"

	"github.com/stressedbypull/swapi-connector/internal/adapters/http/validation"
	"github.com/stressedbypull/swapi-connector/internal/domain"
)

// MapPersonDTOToDomain converts a SWAPI PersonDTO into domain.Person.
func MapPersonDTOToDomain(dto PersonDTO) domain.Person {
	created, err := time.Parse(time.RFC3339, dto.Created)
	if err != nil {
		log.Printf("warn: failed to parse created date '%s': %v", dto.Created, err)
		created = time.Time{}
	}

	mass := validation.ParseMass(dto.Mass)

	return domain.Person{
		Name:   dto.Name,
		Mass:   mass,
		Create: created,
		Films:  dto.Films,
	}
}

// MapPeopleToDomain converts a slice of PersonDTOs to domain.Person slice.
func MapPeopleToDomain(dtos []PersonDTO) []domain.Person {
	if len(dtos) == 0 {
		return []domain.Person{}
	}

	people := make([]domain.Person, 0, len(dtos))
	for _, dto := range dtos {
		people = append(people, MapPersonDTOToDomain(dto))
	}

	return people
}

// MapPlanetDTOToDomain converts a SWAPI PlanetDTO into domain.Planet.
func MapPlanetDTOToDomain(dto PlanetDTO) domain.Planet {
	created, err := time.Parse(time.RFC3339, dto.Created)
	if err != nil {
		log.Printf("warn: failed to parse created date '%s': %v", dto.Created, err)
		created = time.Time{}
	}

	return domain.Planet{
		Name:     dto.Name,
		Resident: dto.Residents,
		Created:  created,
		Films:    dto.Films,
	}
}

// MapPlanetsToDomain converts a slice of PlanetDTOs to domain.Planet slice.
func MapPlanetsToDomain(dtos []PlanetDTO) []domain.Planet {
	if len(dtos) == 0 {
		return []domain.Planet{}
	}

	planets := make([]domain.Planet, 0, len(dtos))
	for _, dto := range dtos {
		planets = append(planets, MapPlanetDTOToDomain(dto))
	}

	return planets
}
