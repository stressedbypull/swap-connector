package sorting

import "sort"

// ByCreated sorts any Sortable entities by created date field.
type ByCreated[T Sortable] struct{}

// Sort sorts entities by created date in ascending or descending order.
func (s ByCreated[T]) Sort(items []T, ascending bool) {
	sort.Slice(items, func(i, j int) bool {
		if ascending {
			return items[i].GetCreated().Before(items[j].GetCreated())
		}
		return items[i].GetCreated().After(items[j].GetCreated())
	})
}
