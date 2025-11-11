package pagination

import "github.com/stressedbypull/swapi-connector/internal/domain"

// BuildResponse creates a paginated response from items and total count
func BuildResponse[T any](items []T, totalCount, page int) domain.PaginatedResponse[T] {
	return domain.PaginatedResponse[T]{
		Count:    totalCount,
		Page:     page,
		PageSize: len(items),
		Results:  items,
	}
}

// TakeFirst returns the first N items from a slice
func TakeFirst[T any](items []T, count int) []T {
	if len(items) <= count {
		return items
	}
	return items[:count]
}
