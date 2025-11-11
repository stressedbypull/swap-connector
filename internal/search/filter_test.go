package search

import (
	"testing"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

func TestFilterPeopleByName(t *testing.T) {
	people := []domain.Person{
		{Name: "Luke Skywalker"},
		{Name: "Darth Vader"},
		{Name: "Anakin Skywalker"},
		{Name: "Obi-Wan Kenobi"},
		{Name: "Leia Organa"},
	}

	tests := []struct {
		name      string
		search    string
		wantNames []string
		wantCount int
	}{
		{
			name:      "search 'sky' returns Skywalkers",
			search:    "sky",
			wantNames: []string{"Luke Skywalker", "Anakin Skywalker"},
			wantCount: 2,
		},
		{
			name:      "search 'vader' returns Darth Vader",
			search:    "vader",
			wantNames: []string{"Darth Vader"},
			wantCount: 1,
		},
		{
			name:      "case insensitive search",
			search:    "SKY",
			wantNames: []string{"Luke Skywalker", "Anakin Skywalker"},
			wantCount: 2,
		},
		{
			name:      "partial match 'an' returns multiple",
			search:    "an",
			wantNames: []string{"Anakin Skywalker", "Obi-Wan Kenobi", "Leia Organa"},
			wantCount: 3,
		},
		{
			name:      "no match returns empty",
			search:    "yoda",
			wantNames: []string{},
			wantCount: 0,
		},
		{
			name:      "empty search returns all",
			search:    "",
			wantNames: []string{"Luke Skywalker", "Darth Vader", "Anakin Skywalker", "Obi-Wan Kenobi", "Leia Organa"},
			wantCount: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterPeopleByName(people, tt.search)

			if len(result) != tt.wantCount {
				t.Errorf("got %d results, want %d", len(result), tt.wantCount)
			}

			for i, person := range result {
				if i >= len(tt.wantNames) {
					t.Errorf("unexpected result at position %d: %s", i, person.Name)
					continue
				}
				if person.Name != tt.wantNames[i] {
					t.Errorf("position %d: got %s, want %s", i, person.Name, tt.wantNames[i])
				}
			}
		})
	}
}

func TestFilterPlanetsByName(t *testing.T) {
	planets := []domain.Planet{
		{Name: "Tatooine"},
		{Name: "Alderaan"},
		{Name: "Hoth"},
	}

	tests := []struct {
		name      string
		search    string
		wantCount int
	}{
		{
			name:      "search 'tato' returns Tatooine",
			search:    "tato",
			wantCount: 1,
		},
		{
			name:      "empty search returns all",
			search:    "",
			wantCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterPlanetsByName(planets, tt.search)

			if len(result) != tt.wantCount {
				t.Errorf("got %d results, want %d", len(result), tt.wantCount)
			}
		})
	}
}
