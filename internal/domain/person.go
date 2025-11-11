package domain

import "time"

type Person struct {
	Name   string    `json:"name"`
	Mass   int       `json:"mass"`
	Create time.Time `json:"created"`
	Films  []string  `json:"films"`
}

// GetName returns the person's name (implements sorting.Sortable).
func (p Person) GetName() string {
	return p.Name
}

// GetCreated returns the creation time (implements sorting.Sortable).
func (p Person) GetCreated() time.Time {
	return p.Create
}
