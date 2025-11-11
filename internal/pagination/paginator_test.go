package pagination_test

import (
	"testing"

	"github.com/stressedbypull/swapi-connector/internal/pagination"
	"github.com/stretchr/testify/assert"
)

func TestAggregationStrategy_CalculatePageRange(t *testing.T) {
	tests := []struct {
		name             string
		requestedPage    int
		desiredPageSize  int
		externalPageSize int
		wantStartPage    int
		wantEndPage      int
		wantPagesNeeded  int
	}{
		{
			name:             "Page 1 with 15 items (external has 10)",
			requestedPage:    1,
			desiredPageSize:  15,
			externalPageSize: 10,
			wantStartPage:    1,
			wantEndPage:      2,
			wantPagesNeeded:  2,
		},
		{
			name:             "Page 2 with 15 items (external has 10)",
			requestedPage:    2,
			desiredPageSize:  15,
			externalPageSize: 10,
			wantStartPage:    2,
			wantEndPage:      3,
			wantPagesNeeded:  2,
		},
		{
			name:             "Page 1 with 10 items (same as external)",
			requestedPage:    1,
			desiredPageSize:  10,
			externalPageSize: 10,
			wantStartPage:    1,
			wantEndPage:      1,
			wantPagesNeeded:  1,
		},
		{
			name:             "Page 3 with 20 items (external has 10)",
			requestedPage:    3,
			desiredPageSize:  20,
			externalPageSize: 10,
			wantStartPage:    5,
			wantEndPage:      6,
			wantPagesNeeded:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strategy := pagination.NewAggregationStrategy(
				tt.requestedPage,
				tt.desiredPageSize,
				tt.externalPageSize,
			)

			gotStartPage, gotEndPage, gotPagesNeeded := strategy.CalculatePageRange()

			assert.Equal(t, tt.wantStartPage, gotStartPage, "startPage mismatch")
			assert.Equal(t, tt.wantEndPage, gotEndPage, "endPage mismatch")
			assert.Equal(t, tt.wantPagesNeeded, gotPagesNeeded, "pagesNeeded mismatch")
		})
	}
}

func TestAggregationStrategy_CalculateOffset(t *testing.T) {
	tests := []struct {
		name             string
		requestedPage    int
		desiredPageSize  int
		externalPageSize int
		startPage        int
		wantOffset       int
	}{
		{
			name:             "Page 1 starts at SWAPI page 1, offset 0",
			requestedPage:    1,
			desiredPageSize:  15,
			externalPageSize: 10,
			startPage:        1,
			wantOffset:       0,
		},
		{
			name:             "Page 2 starts at SWAPI page 2, offset 5",
			requestedPage:    2,
			desiredPageSize:  15,
			externalPageSize: 10,
			startPage:        2,
			wantOffset:       5,
		},
		{
			name:             "Page 3 starts at SWAPI page 3, offset 10",
			requestedPage:    3,
			desiredPageSize:  15,
			externalPageSize: 10,
			startPage:        3,
			wantOffset:       10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strategy := pagination.NewAggregationStrategy(
				tt.requestedPage,
				tt.desiredPageSize,
				tt.externalPageSize,
			)

			gotOffset := strategy.CalculateOffset(tt.startPage)
			assert.Equal(t, tt.wantOffset, gotOffset)
		})
	}
}

func TestSliceResultsTyped(t *testing.T) {
	tests := []struct {
		name             string
		fetchedItems     []string // The items actually fetched from external API
		requestedPage    int
		desiredPageSize  int
		externalPageSize int
		wantResults      []string
	}{
		{
			name: "Page 1: fetch pages 1-2 (20 items), return first 15",
			fetchedItems: []string{
				"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", // Page 1
				"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", // Page 2
			},
			requestedPage:    1,
			desiredPageSize:  15,
			externalPageSize: 10,
			wantResults:      []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"},
		},
		{
			name: "Page 2: fetch pages 2-3 (20 items), offset 5, return 15",
			fetchedItems: []string{
				"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", // Page 2 (items 10-19)
				"U", "V", "W", "X", "Y", "Z", "AA", "BB", "CC", "DD", // Page 3 (items 20-29)
			},
			requestedPage:    2,
			desiredPageSize:  15,
			externalPageSize: 10,
			wantResults:      []string{"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "BB", "CC", "DD"},
		},
		{
			name: "Page 3: fetch pages 5-6 (15 items available), offset 0, return 10",
			fetchedItems: []string{
				"EE", "FF", "GG", "HH", "II", "JJ", "KK", "LL", "MM", "NN", // Page 5 (items 40-49)
			},
			requestedPage:    3,
			desiredPageSize:  15,
			externalPageSize: 10,
			wantResults:      []string{"EE", "FF", "GG", "HH", "II", "JJ", "KK", "LL", "MM", "NN"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strategy := pagination.NewAggregationStrategy(
				tt.requestedPage,
				tt.desiredPageSize,
				tt.externalPageSize,
			)

			results := pagination.SliceResultsTyped(tt.fetchedItems, strategy)
			assert.Equal(t, tt.wantResults, results)
		})
	}
}
