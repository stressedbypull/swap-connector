package handlers

import "time"

// Person represents a Star Wars character.
//
// @Description Star Wars character information
type Person struct {
	Name    string    `json:"name" example:"Luke Skywalker"`
	Mass    int       `json:"mass" example:"77"`
	Created time.Time `json:"created" example:"2014-12-09T13:50:51.644000Z"`
	Films   []string  `json:"films" example:"https://swapi.dev/api/films/1/,https://swapi.dev/api/films/2/"`
}

// PeopleListResponse represents a paginated response of people.
//
// @Description Paginated list of Star Wars characters
type PeopleListResponse struct {
	Count    int      `json:"count" example:"82"`
	Page     int      `json:"page" example:"1"`
	PageSize int      `json:"pageSize" example:"15"`
	Results  []Person `json:"results"`
}

// ErrorDetail contains error information.
//
// @Description Detailed error information
type ErrorDetail struct {
	Message string                 `json:"message" example:"Person not found"`
	Code    string                 `json:"code,omitempty" example:"NOT_FOUND"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// ErrorResponse represents an error API response.
//
// @Description API error response
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}
