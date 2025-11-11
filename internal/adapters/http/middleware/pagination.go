package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/pagination"
)

const (
	DefaultPage = 1
	PageSize    = 15 // Fixed per requirements
)

type PaginationParams struct {
	Page     int
	PageSize int
}

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse page (only parameter we accept)
		page, err := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(DefaultPage)))
		if err != nil || page < 1 {
			page = DefaultPage
		}

		c.Set("pagination", PaginationParams{
			Page:     page,
			PageSize: PageSize, // Always 15
		})

		c.Next()
	}
}

func GetPaginationParams(c *gin.Context) PaginationParams {
	if params, exists := c.Get("pagination"); exists {
		return params.(PaginationParams)
	}
	return PaginationParams{
		Page:     DefaultPage,
		PageSize: PageSize,
	}
}

// aggregateAndBuildResponse combines multiple pages and builds final response
func aggregateAndBuildResponse(pages []PageData, page int) domain.PaginatedResponse[domain.Person] {
	// Aggregate all people
	var allPeople []domain.Person
	var totalCount int

	for _, p := range pages {
		allPeople = append(allPeople, p.People...)
		totalCount = p.Count
	}

	// Slice to get exact page
	results := pagination.SliceResults(allPeople, page, PageSize)

	return domain.PaginatedResponse[domain.Person]{
		Count:    totalCount,
		Page:     page,
		PageSize: len(results),
		Results:  results,
	}
}
