package sorting

import (
	"testing"
	"time"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

func TestByCreated_Sort(t *testing.T) {
	date1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	date3 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name      string
		people    []domain.Person
		ascending bool
		wantDates []time.Time
	}{
		{
			name: "sort ascending",
			people: []domain.Person{
				{Name: "Person2", Create: date2},
				{Name: "Person3", Create: date3},
				{Name: "Person1", Create: date1},
			},
			ascending: true,
			wantDates: []time.Time{date1, date2, date3},
		},
		{
			name: "sort descending",
			people: []domain.Person{
				{Name: "Person2", Create: date2},
				{Name: "Person1", Create: date1},
				{Name: "Person3", Create: date3},
			},
			ascending: false,
			wantDates: []time.Time{date3, date2, date1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter := ByCreated[domain.Person]{}
			sorter.Sort(tt.people, tt.ascending)

			for i, person := range tt.people {
				if !person.Create.Equal(tt.wantDates[i]) {
					t.Errorf("position %d: got %v, want %v", i, person.Create, tt.wantDates[i])
				}
			}
		})
	}
}
