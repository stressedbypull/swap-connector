package swapi

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to load test fixture files
func loadTestFixture(t *testing.T, filename string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", filename))
	require.NoError(t, err, "Failed to read test fixture: %s", filename)
	return data
}

func TestClient_ParsePeopleResponse(t *testing.T) {
	t.Run("Parse SWAPI people response from fixture", func(t *testing.T) {
		// Given: load fixture
		mockData := loadTestFixture(t, "people_response.json")

		// When: parse the SWAPI response
		var swapiResp SWAPIPeopleResponse
		err := json.Unmarshal(mockData, &swapiResp)
		require.NoError(t, err)

		// Then: verify structure
		assert.Equal(t, 82, swapiResp.Count)
		assert.NotNil(t, swapiResp.Next)
		assert.Nil(t, swapiResp.Previous)
		assert.Len(t, swapiResp.Results, 3)

		// Map to domain
		people := make([]domain.Person, 0, len(swapiResp.Results))
		for _, dto := range swapiResp.Results {
			people = append(people, MapPersonDTOToDomain(dto))
		}

		// Verify domain mapping
		assert.Equal(t, "Luke Skywalker", people[0].Name)
		assert.Equal(t, 77, people[0].Mass)
		assert.False(t, people[0].Create.IsZero())

		assert.Equal(t, "Darth Vader", people[1].Name)
		assert.Equal(t, 136, people[1].Mass)

		assert.Equal(t, "Leia Organa", people[2].Name)
		assert.Equal(t, 49, people[2].Mass)
	})
}
