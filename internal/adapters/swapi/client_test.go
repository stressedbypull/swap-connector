package swapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_FetchAllPeople(t *testing.T) {
	url := "https://swapi.dev/api"
	client := NewClient(url)
	t.Run("Get all people from swapi", func(t *testing.T) {
		ctx := context.Background()
		result, err := client.FetchPeople(ctx, 1, 15, "")

		assert.NoError(t, err)
		assert.Equal(t, 0, result.Count)
		assert.Empty(t, result.Results)
	})
}
