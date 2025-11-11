package sorting

import (
	"testing"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

func TestByName_Sort(t *testing.T) {
	tests := []struct {
		name      string
		people    []domain.Person
		ascending bool
		want      []string
	}{
		{
			name: "sort ascending",
			people: []domain.Person{
				{Name: "Luke Skywalker"},
				{Name: "Darth Vader"},
				{Name: "Anakin Skywalker"},
			},
			ascending: true,
			want:      []string{"Anakin Skywalker", "Darth Vader", "Luke Skywalker"},
		},
		{
			name: "sort descending",
			people: []domain.Person{
				{Name: "Luke Skywalker"},
				{Name: "Darth Vader"},
				{Name: "Anakin Skywalker"},
			},
			ascending: false,
			want:      []string{"Luke Skywalker", "Darth Vader", "Anakin Skywalker"},
		},
		{
			name: "case insensitive",
			people: []domain.Person{
				{Name: "luke skywalker"},
				{Name: "DARTH VADER"},
				{Name: "Anakin Skywalker"},
			},
			ascending: true,
			want:      []string{"Anakin Skywalker", "DARTH VADER", "luke skywalker"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter := ByName[domain.Person]{}
			sorter.Sort(tt.people, tt.ascending)

			for i, person := range tt.people {
				if person.Name != tt.want[i] {
					t.Errorf("position %d: got %s, want %s", i, person.Name, tt.want[i])
				}
			}
		})
	}
}
