package sorting

import (
	"testing"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

func TestByCreated_Sort(t *testing.T) {
	date1 := "2020-01-01"
	date2 := "2021-01-01"
	date3 := "2022-01-01"

	tests := []struct {
		name      string
		people    []domain.Person
		ascending bool
		wantDates []string
	}{
		{
			name: "sort ascending",
			people: []domain.Person{
				{Name: "Person2", Create: date2},
				{Name: "Person3", Create: date3},
				{Name: "Person1", Create: date1},
			},
			ascending: true,
			wantDates: []string{date1, date2, date3},
		},
		{
			name: "sort descending",
			people: []domain.Person{
				{Name: "Person2", Create: date2},
				{Name: "Person1", Create: date1},
				{Name: "Person3", Create: date3},
			},
			ascending: false,
			wantDates: []string{date3, date2, date1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter := ByCreated[domain.Person]{}
			sorter.Sort(tt.people, tt.ascending)

			for i, person := range tt.people {
				if person.Create != tt.wantDates[i] {
					t.Errorf("position %d: got %v, want %v", i, person.Create, tt.wantDates[i])
				}
			}
		})
	}
}
