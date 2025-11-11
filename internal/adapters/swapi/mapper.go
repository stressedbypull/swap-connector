package swapi

import (
	"log"
	"strconv"
	"time"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/pagination"
)

// MapPersonDTOToDomain converts a SWAPI PersonDTO into domain.Person
func MapPersonDTOToDomain(p PersonDTO) domain.Person {
	created, err := time.Parse(time.RFC3339, p.Created)
	if err != nil {
		log.Printf("warn: failed to parse created date '%s': %v", p.Created, err)
		created = time.Time{}
	}

	mass := ParseMass(p.Mass)

	return domain.Person{
		Name:   p.Name,
		Mass:   mass,
		Create: created,
		Films:  p.Films,
	}
}

// MapPeopleToDomain converts a slice of PersonDTOs to domain.Person slice
func MapPeopleToDomain(dtos []PersonDTO) []domain.Person {
	people := make([]domain.Person, 0, len(dtos))
	for _, dto := range dtos {
		people = append(people, MapPersonDTOToDomain(dto))
	}
	return people
}

// HasNextPage checks if there's a next page in SWAPI response
func HasNextPage(next *string) bool {
	return next != nil
}

// MapSwapiPeopleResponseToPageData converts SWAPI people response to PageData
func MapSwapiPeopleResponseToPageData(swapiResp *SWAPIPeopleResponse) pagination.PageData[domain.Person] {
	var singlePage pagination.PageData[domain.Person]

	singlePage.Items = MapPeopleToDomain(swapiResp.Results)
	singlePage.Count = swapiResp.Count
	singlePage.HasNext = HasNextPage(swapiResp.Next)

	return singlePage
}

// MapSwapiPlanetsResponseToPageData converts SWAPI planets response to PageData
func MapSwapiPlanetsResponseToPageData(swapiResp *SWAPIPlanetsResponse) pagination.PageData[domain.Planet] {
	var singlePage pagination.PageData[domain.Planet]

	singlePage.Items = MapPlanetsToDomain(swapiResp.Results)
	singlePage.Count = swapiResp.Count
	singlePage.HasNext = HasNextPage(swapiResp.Next)

	return singlePage
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

// MapPlanetsToDomain converts a slice of PlanetDTOs to domain.Planet slice
func MapPlanetsToDomain(dtos []PlanetDTO) []domain.Planet {
	planets := make([]domain.Planet, 0, len(dtos))
	for _, dto := range dtos {
		planets = append(planets, MapPlanetDTOToDomain(dto))
	}
	return planets
}

// ParseMass converts mass string to int, handling "unknown" values
func ParseMass(s string) int {
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
