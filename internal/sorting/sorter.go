package sorting

import "time"

// Sortable defines entities that can be sorted by common fields.
type Sortable interface {
	GetName() string
	GetCreated() time.Time
}

// Sorter defines the interface for sorting strategies (Open-Closed Principle).
// New sorting strategies can be added without modifying existing code.
// Generic interface works with any Sortable type.
type Sorter[T Sortable] interface {
	Sort(items []T, ascending bool)
}
