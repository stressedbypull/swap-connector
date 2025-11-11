package domain

import "time"

type Planet struct {
	Name     string   `json:"name"`
	Resident []string `json:"residents"`
	Created  string   `json:"created"`
	Films    []string `json:"films"`
}

// GetName returns the planet's name (implements sorting.Sortable).
func (p Planet) GetName() string {
	return p.Name
}

// GetCreated returns the creation time (implements sorting.Sortable).
func (p Planet) GetCreated() time.Time {
	// Parse the date string for sorting purposes
	t, _ := time.Parse("2006-01-02", p.Created)
	return t
}
