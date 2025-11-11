package sorting

import (
	"testing"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

func TestByMass_Sort(t *testing.T) {
	tests := []struct {
		name      string
		people    []domain.Person
		ascending bool
		wantMass  []int
	}{
		{
			name: "sort ascending",
			people: []domain.Person{
				{Name: "Person2", Mass: 80},
				{Name: "Person3", Mass: 136},
				{Name: "Person1", Mass: 77},
			},
			ascending: true,
			wantMass:  []int{77, 80, 136},
		},
		{
			name: "sort descending",
			people: []domain.Person{
				{Name: "Person2", Mass: 80},
				{Name: "Person1", Mass: 77},
				{Name: "Person3", Mass: 136},
			},
			ascending: false,
			wantMass:  []int{136, 80, 77},
		},
		{
			name: "with zero mass",
			people: []domain.Person{
				{Name: "Person2", Mass: 80},
				{Name: "Person3", Mass: 0},
				{Name: "Person1", Mass: 77},
			},
			ascending: true,
			wantMass:  []int{0, 77, 80},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter := ByMass{}
			sorter.Sort(tt.people, tt.ascending)

			for i, person := range tt.people {
				if person.Mass != tt.wantMass[i] {
					t.Errorf("position %d: got %d, want %d", i, person.Mass, tt.wantMass[i])
				}
			}
		})
	}
}
