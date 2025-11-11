package search

import (
	"strings"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

// FilterPeopleByName filters people by name using case-insensitive partial match.
// Example: "sky" matches "Luke Skywalker", "Anakin Skywalker".
func FilterPeopleByName(people []domain.Person, search string) []domain.Person {
	if search == "" {
		return people
	}

	searchLower := strings.ToLower(search)
	filtered := make([]domain.Person, 0)

	for _, person := range people {
		if strings.Contains(strings.ToLower(person.Name), searchLower) {
			filtered = append(filtered, person)
		}
	}

	return filtered
}

// FilterPlanetsByName filters planets by name using case-insensitive partial match.
func FilterPlanetsByName(planets []domain.Planet, search string) []domain.Planet {
	if search == "" {
		return planets
	}

	searchLower := strings.ToLower(search)
	filtered := make([]domain.Planet, 0)

	for _, planet := range planets {
		if strings.Contains(strings.ToLower(planet.Name), searchLower) {
			filtered = append(filtered, planet)
		}
	}

	return filtered
}
