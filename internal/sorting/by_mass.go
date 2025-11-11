package sorting

import (
	"sort"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

// ByMass sorts people by mass field (Person-specific sorter).
type ByMass struct{}

// Sort sorts people by mass in ascending or descending order.
// Note: This is Person-specific and doesn't use generics since mass is unique to Person.
func (s ByMass) Sort(people []domain.Person, ascending bool) {
	sort.Slice(people, func(i, j int) bool {
		if ascending {
			return people[i].Mass < people[j].Mass
		}
		return people[i].Mass > people[j].Mass
	})
}
