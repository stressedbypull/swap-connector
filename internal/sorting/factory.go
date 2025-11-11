package sorting

import "github.com/stressedbypull/swapi-connector/internal/domain"

// NewPersonSorter creates a sorter for Person entities based on the field name.
// Returns nil if the field is not supported.
func NewPersonSorter(field string) Sorter[domain.Person] {
	switch field {
	case "name":
		return ByName[domain.Person]{}
	case "created":
		return ByCreated[domain.Person]{}
	case "mass":
		// Mass is Person-specific, but we wrap it to match the interface
		return personMassAdapter{}
	default:
		return nil
	}
}

// NewPlanetSorter creates a sorter for Planet entities based on the field name.
// Returns nil if the field is not supported.
func NewPlanetSorter(field string) Sorter[domain.Planet] {
	switch field {
	case "name":
		return ByName[domain.Planet]{}
	case "created":
		return ByCreated[domain.Planet]{}
	default:
		return nil
	}
}

// personMassAdapter adapts ByMass to the generic Sorter interface.
type personMassAdapter struct{}

func (a personMassAdapter) Sort(people []domain.Person, ascending bool) {
	ByMass{}.Sort(people, ascending)
}
