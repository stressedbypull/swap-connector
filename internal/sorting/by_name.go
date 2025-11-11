package sorting

import (
	"sort"
	"strings"
)

// ByName sorts any Sortable entities by name field.
type ByName[T Sortable] struct{}

// Sort sorts entities by name in ascending or descending order.
func (s ByName[T]) Sort(items []T, ascending bool) {
	sort.Slice(items, func(i, j int) bool {
		cmp := strings.Compare(
			strings.ToLower(items[i].GetName()),
			strings.ToLower(items[j].GetName()),
		)
		if ascending {
			return cmp < 0
		}
		return cmp > 0
	})
}
