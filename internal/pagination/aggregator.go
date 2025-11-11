package pagination

// PageData represents data fetched from a single external API page
// Generic type to work with any domain entity
type PageData[T any] struct {
	Items   []T
	Count   int
	HasNext bool
}

// Aggregate combines multiple PageData into one flat list
func Aggregate[T any](pages []PageData[T]) ([]T, int) {
	var allItems []T
	var totalCount int

	for _, page := range pages {
		allItems = append(allItems, page.Items...)
		totalCount = page.Count // Same for all pages
	}

	return allItems, totalCount
}
