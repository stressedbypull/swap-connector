package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/middleware"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/response"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/validation"
)

// Allowed values for people endpoint
var (
	allowedPeopleSortBy = []string{"name", "created", "mass"} // Fields that can be sorted
	allowedSortOrder    = []string{"asc", "desc"}             // Sort directions
)

// PeopleQueryParams holds the validated query parameters for the people endpoint.
type PeopleQueryParams struct {
	Page      int    // Which page (from pagination middleware)
	Search    string // Text to search in names (optional)
	SortBy    string // Field to sort by: "name", "created", or "mass" (optional)
	SortOrder string // Sort direction: "asc" or "desc" (default: "asc")
}

// ParsePeopleQueryParams gets query parameters from middleware and validates them.
//
// Flow:
//  1. Get page number (already validated by pagination middleware)
//  2. Get search/sort values (from query middleware)
//  3. Validate sortBy is one of: name, created, mass
//  4. Validate sortOrder is one of: asc, desc
//  5. Return validated params OR send error response
//
// Returns:
//   - PeopleQueryParams: the validated parameters
//   - bool: true if valid, false if validation failed (error already sent to client)
func ParsePeopleQueryParams(c *gin.Context) (PeopleQueryParams, bool) {
	// Step 1: Get page number (middleware already checked it's >= 1)
	paginationParams := middleware.GetPaginationParams(c)

	// Step 2: Get search and sort parameters (not validated yet)
	queryParams := middleware.GetQueryParams(c)

	// Step 3: Validate the parameters
	validator := validation.New()

	// Only validate sortBy if user provided it
	if queryParams.SortBy != "" {
		validator.ValidateOneOf("sortBy", queryParams.SortBy, allowedPeopleSortBy)
	}

	// Always validate sortOrder (has default "asc")
	validator.ValidateOneOf("sortOrder", queryParams.SortOrder, allowedSortOrder)

	// Step 4: If validation failed, send error to client and return false
	if validator.HasErrors() {
		response.ValidationError(c, validator.ErrorsMap())
		return PeopleQueryParams{}, false
	}

	// Step 5: All good! Return the validated parameters
	return PeopleQueryParams{
		Page:      paginationParams.Page,
		Search:    queryParams.Search,
		SortBy:    queryParams.SortBy,
		SortOrder: queryParams.SortOrder,
	}, true
}
