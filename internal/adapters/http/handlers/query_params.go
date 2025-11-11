package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	// Pagination defaults
	DefaultPage = 1
	MinPage     = 1

	// Sort order values
	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"
)

// PeopleQueryParams represents query parameters for people endpoint.
type PeopleQueryParams struct {
	Page      int
	Search    string
	SortBy    string
	SortOrder string
}

// ParsePeopleQueryParams extracts and validates query parameters from the request.
func ParsePeopleQueryParams(c *gin.Context) PeopleQueryParams {
	params := PeopleQueryParams{
		Page:      DefaultPage,
		Search:    "",
		SortBy:    "",
		SortOrder: SortOrderAsc,
	}

	// Parse page number
	if pageStr := c.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page >= MinPage {
			params.Page = page
		}
	}

	// Get search parameter
	params.Search = c.Query("search")

	// Get sort parameters
	params.SortBy = c.Query("sortBy")

	// Validate sort order
	if sortOrder := c.Query("sortOrder"); sortOrder != "" {
		if sortOrder == SortOrderAsc || sortOrder == SortOrderDesc {
			params.SortOrder = sortOrder
		}
	}

	return params
}
