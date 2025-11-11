package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/middleware"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/response"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/validation"
)

var (
	// Allowed sort fields for people endpoint
	allowedPeopleSortBy = []string{"name", "created", "mass"}
	// Allowed sort orders
	allowedSortOrder = []string{"asc", "desc"}
)

// PeopleQueryParams represents validated query parameters for people endpoint.
type PeopleQueryParams struct {
	Page      int
	Search    string
	SortBy    string
	SortOrder string
}

// ParsePeopleQueryParams extracts and validates query parameters from the request.
// Returns false if validation fails (response already sent).
func ParsePeopleQueryParams(c *gin.Context) (PeopleQueryParams, bool) {
	// Get page from pagination middleware (already validated)
	paginationParams := middleware.GetPaginationParams(c)

	// Get search/sort from query middleware
	queryParams := middleware.GetQueryParams(c)

	// Create validator
	validator := validation.New()

	// Validate sortBy if provided
	if queryParams.SortBy != "" {
		validator.ValidateOneOf("sortBy", queryParams.SortBy, allowedPeopleSortBy)
	}

	// Validate sortOrder if provided
	validator.ValidateOneOf("sortOrder", queryParams.SortOrder, allowedSortOrder)

	// If validation failed, return error response
	if validator.HasErrors() {
		response.ValidationError(c, validator.ErrorsMap())
		return PeopleQueryParams{}, false
	}

	// Return validated params
	return PeopleQueryParams{
		Page:      paginationParams.Page,
		Search:    queryParams.Search,
		SortBy:    queryParams.SortBy,
		SortOrder: queryParams.SortOrder,
	}, true
}
