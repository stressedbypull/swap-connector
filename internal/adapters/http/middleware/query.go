package middleware

import (
	"github.com/gin-gonic/gin"
)

// QueryParams holds common query parameters for search, sort, and filter.
type QueryParams struct {
	Search    string
	SortBy    string
	SortOrder string
}

const queryParamsKey = "queryParams"

// QueryMiddleware extracts common query parameters (search, sortBy, sortOrder)
// and stores them in the context for use by handlers.
func QueryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := QueryParams{
			Search:    c.DefaultQuery("search", ""),
			SortBy:    c.DefaultQuery("sortBy", ""),
			SortOrder: c.DefaultQuery("sortOrder", "asc"),
		}

		c.Set(queryParamsKey, params)
		c.Next()
	}
}

// GetQueryParams retrieves QueryParams from context.
func GetQueryParams(c *gin.Context) QueryParams {
	params, exists := c.Get(queryParamsKey)
	if !exists {
		return QueryParams{SortOrder: "asc"}
	}
	return params.(QueryParams)
}
