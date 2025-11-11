package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/middleware"
)

const (
	// Sort order values
	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"
)

// PeopleQueryParams represents query parameters for people endpoint.
type PeopleQueryParams struct {
	Page      int    `form:"page"`
	Search    string `form:"search"`
	SortBy    string `form:"sortBy" binding:"omitempty,oneof=name created mass"`
	SortOrder string `form:"sortOrder" binding:"omitempty,oneof=asc desc"`
}

// ParsePeopleQueryParams extracts and validates query parameters from the request.
// Uses middleware for page validation and Gin's binding for other params.
func ParsePeopleQueryParams(c *gin.Context) PeopleQueryParams {
	// Get page from middleware (already validated)
	paginationParams := middleware.GetPaginationParams(c)

	// Bind and validate other query params
	var query PeopleQueryParams
	if err := c.ShouldBindQuery(&query); err != nil {
		// If validation fails, use defaults
		query = PeopleQueryParams{
			SortOrder: SortOrderAsc,
		}
	}

	// Override page with middleware value
	query.Page = paginationParams.Page

	// Set default sort order if empty
	if query.SortOrder == "" {
		query.SortOrder = SortOrderAsc
	}

	return query
}
