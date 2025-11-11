package search

import (
	"strings"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/errors"
)

// FilterPeopleByName filters people by name using case-insensitive partial match.
// Example: "sky" matches "Luke Skywalker", "Anakin Skywalker".
// Returns ErrPersonNotFound if search is provided but no results are found.
func FilterPeopleByName(people []domain.Person, search string) ([]domain.Person, error) {
	if search == "" {
		return people, nil
	}

	searchLower := strings.ToLower(search)
	filtered := make([]domain.Person, 0)

	for _, person := range people {
		if strings.Contains(strings.ToLower(person.Name), searchLower) {
			filtered = append(filtered, person)
		}
	}

	if len(filtered) == 0 {
		return nil, errors.ErrPersonNotFound
	}

	return filtered, nil
}

// FilterPlanetsByName filters planets by name using case-insensitive partial match.
// Returns ErrPlanetNotFound if search is provided but no results are found.
func FilterPlanetsByName(planets []domain.Planet, search string) ([]domain.Planet, error) {
	if search == "" {
		return planets, nil
	}

	searchLower := strings.ToLower(search)
	filtered := make([]domain.Planet, 0)

	for _, planet := range planets {
		if strings.Contains(strings.ToLower(planet.Name), searchLower) {
			filtered = append(filtered, planet)
		}
	}

	if len(filtered) == 0 {
		return nil, errors.ErrPlanetNotFound
	}

	return filtered, nil
}
