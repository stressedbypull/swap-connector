package pagination

import "github.com/stressedbypull/swapi-connector/internal/domain"

// AggregationStrategy represents the strategy for aggregating multiple external API pages
// into a single page with custom page size.
type AggregationStrategy struct {
	RequestedPage    int // The page number the client requested
	DesiredPageSize  int // The page size we want to return (e.g., 15)
	ExternalPageSize int // The page size of the external API (e.g., 10)
}

// NewAggregationStrategy creates a new pagination aggregation strategy.
func NewAggregationStrategy(requestedPage, desiredPageSize, externalPageSize int) *AggregationStrategy {
	return &AggregationStrategy{
		RequestedPage:    requestedPage,
		DesiredPageSize:  desiredPageSize,
		ExternalPageSize: externalPageSize,
	}
}

// CalculatePageRange determines which external API pages need to be fetched.
// Returns: startPage, endPage, pagesNeeded
//
// Example: Page 1 with 15 items (external has 10/page)
//   - Need items 0-14 → SWAPI pages 1-2
//
// Example: Page 2 with 15 items
//   - Need items 15-29 → SWAPI pages 2-3
func (s *AggregationStrategy) CalculatePageRange() (startPage, endPage, pagesNeeded int) {
	// Calculate the absolute item range we need
	startItem := (s.RequestedPage - 1) * s.DesiredPageSize
	endItem := startItem + s.DesiredPageSize

	// Calculate which external API pages contain these items
	startPage = startItem/s.ExternalPageSize + 1
	endPage = (endItem-1)/s.ExternalPageSize + 1
	pagesNeeded = endPage - startPage + 1

	return startPage, endPage, pagesNeeded
}

// CalculateOffset calculates the offset within the fetched data where our results begin.
//
// Example: If we fetched SWAPI pages 2-3 (items 10-29 from SWAPI)
// and need our page 2 (items 15-29), the offset is 15 - 10 = 5
func (s *AggregationStrategy) CalculateOffset(startPage int) int {
	startItem := (s.RequestedPage - 1) * s.DesiredPageSize
	return startItem - (startPage-1)*s.ExternalPageSize
}

// SliceResults extracts the correct items for the requested page from aggregated data.
func (s *AggregationStrategy) SliceResults(fetchedItems []interface{}) []interface{} {
	if len(fetchedItems) == 0 {
		return []interface{}{}
	}

	startPage, _, _ := s.CalculatePageRange()
	offset := s.CalculateOffset(startPage)

	if offset >= len(fetchedItems) {
		return []interface{}{}
	}

	// Slice to get exactly DesiredPageSize items (or remaining items)
	end := offset + s.DesiredPageSize
	if end > len(fetchedItems) {
		end = len(fetchedItems)
	}

	return fetchedItems[offset:end]
}

// SliceResultsTyped is a generic version of SliceResults that preserves type safety.
func SliceResultsTyped[T any](fetchedItems []T, strategy *AggregationStrategy) []T {
	if len(fetchedItems) == 0 {
		return []T{}
	}

	startPage, _, _ := strategy.CalculatePageRange()
	offset := strategy.CalculateOffset(startPage)

	if offset >= len(fetchedItems) {
		return []T{}
	}

	// Slice to get exactly DesiredPageSize items (or remaining items)
	end := offset + strategy.DesiredPageSize
	if end > len(fetchedItems) {
		end = len(fetchedItems)
	}

	return fetchedItems[offset:end]
}

// BuildResponse creates a paginated response from aggregated data.
func BuildResponse[T any](items []T, totalCount int, strategy *AggregationStrategy) domain.PaginatedResponse[T] {
	results := SliceResultsTyped(items, strategy)

	return domain.PaginatedResponse[T]{
		Count:    totalCount,
		Page:     strategy.RequestedPage,
		PageSize: len(results),
		Results:  results,
	}
}
