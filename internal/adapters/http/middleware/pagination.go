package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DefaultPage = 1
	MinPage     = 1
)

// PaginationParams holds parsed pagination parameters.
type PaginationParams struct {
	Page int
}

// PaginationMiddleware extracts and validates page parameter from query string.
// Sets it in the context for handlers to use.
func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := DefaultPage

		if pageStr := c.Query("page"); pageStr != "" {
			if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage >= MinPage {
				page = parsedPage
			}
		}

		c.Set("pagination", PaginationParams{
			Page: page,
		})

		c.Next()
	}
}

// GetPaginationParams retrieves pagination params from context.
func GetPaginationParams(c *gin.Context) PaginationParams {
	if params, exists := c.Get("pagination"); exists {
		return params.(PaginationParams)
	}

	return PaginationParams{
		Page: DefaultPage,
	}
}
