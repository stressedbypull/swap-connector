package domain

import "time"

// Person represents a Star Wars character
// @name Person
type Person struct {
	Name   string   `json:"name" example:"Luke Skywalker"`
	Mass   int      `json:"mass" example:"77"`
	Create string   `json:"created" example:"2014-12-09"`
	Films  []string `json:"films" example:"https://swapi.dev/api/films/1/,https://swapi.dev/api/films/2/"`
}

// GetName returns the person's name (implements sorting.Sortable).
func (p Person) GetName() string {
	return p.Name
}

// GetCreated returns the creation time (implements sorting.Sortable).
func (p Person) GetCreated() time.Time {
	// Parse the date string for sorting purposes
	t, _ := time.Parse("2006-01-02", p.Create)
	return t
}
