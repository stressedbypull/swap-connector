package swapi

import (
	"testing"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stretchr/testify/assert"
)

// Test the mapper function directly (no HTTP, no mock)
func TestMapPersonDTOToDomain(t *testing.T) {
	// Given: a PersonDTO from SWAPI
	dto := PersonDTO{
		Name:    "Luke Skywalker",
		Mass:    "77",
		Created: "2014-12-09T13:50:51.644000Z",
		Films: []string{
			"https://swapi.dev/api/films/1/",
			"https://swapi.dev/api/films/2/",
		},
	}

	// When: we map it to domain
	result := MapPersonDTOToDomain(dto)

	// Then: it should be correctly converted (date only in YYYY-MM-DD format)
	expected := domain.Person{
		Name:   "Luke Skywalker",
		Mass:   77,
		Create: "2014-12-09",
		Films: []string{
			"https://swapi.dev/api/films/1/",
			"https://swapi.dev/api/films/2/",
		},
	}

	assert.Equal(t, expected.Name, result.Name)
	assert.Equal(t, expected.Mass, result.Mass)
	assert.Equal(t, expected.Create, result.Create)
	assert.Equal(t, expected.Films, result.Films)
}
