package middleware

import (
	"github.com/gin-gonic/gin"
)

// QueryParams holds search and sorting parameters extracted from URL query string.
// Example: ?search=luke&sortBy=name&sortOrder=asc
type QueryParams struct {
	Search    string // Optional: filter by name (e.g., "sky")
	SortBy    string // Optional: field to sort by (e.g., "name", "created")
	SortOrder string // Optional: "asc" or "desc" (default: "asc")
}

// QueryMiddleware extracts search and sort parameters from the URL query string.
// It does NOT validate them - validation happens in the handler for each specific resource.
//
// This middleware just extracts the raw values:
//   - search: whatever the user typed
//   - sortBy: field name (validated later per resource)
//   - sortOrder: defaults to "asc" if not provided
//
// Example URL: /api/people?search=luke&sortBy=name&sortOrder=desc
func QueryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract query parameters from URL
		search := c.Query("search")       // Get "search" param (empty string if not present)
		sortBy := c.Query("sortBy")       // Get "sortBy" param (empty string if not present)
		sortOrder := c.Query("sortOrder") // Get "sortOrder" param (empty string if not present)

		// Set default for sortOrder if empty
		if sortOrder == "" {
			sortOrder = "asc"
		}

		// Store in context so handlers can access it
		c.Set("queryParams", QueryParams{
			Search:    search,
			SortBy:    sortBy,
			SortOrder: sortOrder,
		})

		c.Next() // Continue to next middleware/handler
	}
}

// GetQueryParams retrieves the query parameters from the request context.
// Handlers call this to get search/sort values.
func GetQueryParams(c *gin.Context) QueryParams {
	value, exists := c.Get("queryParams")
	if !exists {
		// Fallback if middleware wasn't called (shouldn't happen)
		return QueryParams{SortOrder: "asc"}
	}

	// Safe type assertion with check
	params, ok := value.(QueryParams)
	if !ok {
		// Defensive: return default if type assertion fails
		return QueryParams{SortOrder: "asc"}
	}

	return params
}
