package integrational

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/handlers"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/middleware"
	"github.com/stressedbypull/swapi-connector/internal/adapters/swapi"
	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPeopleHandler_Integration tests against REAL SWAPI
func TestPeopleHandler_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	gin.SetMode(gin.TestMode)

	// Setup REAL dependencies - connecting to actual SWAPI
	httpClient := &http.Client{Timeout: 10 * time.Second}
	swapiClient := swapi.NewClient("https://swapi.dev/api", httpClient)
	peopleService := services.NewPeopleService(swapiClient)
	handler := handlers.NewPeopleHandler(peopleService)

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		checkResponse  func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name:           "fetch real people from SWAPI",
			url:            "/people",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var resp domain.PaginatedResponse[domain.Person]
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				require.NoError(t, err)

				// Real SWAPI has 82 people total
				assert.Greater(t, resp.Count, 0, "Should return people from real SWAPI")
				assert.Equal(t, 1, resp.Page)
				assert.NotEmpty(t, resp.Results, "Should have results")

				// Check that we got real data
				if len(resp.Results) > 0 {
					assert.NotEmpty(t, resp.Results[0].Name)
					assert.Greater(t, resp.Results[0].Mass, 0)
				}
			},
		},
		{
			name:           "search for Luke Skywalker",
			url:            "/people?search=Luke",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var resp domain.PaginatedResponse[domain.Person]
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				require.NoError(t, err)

				assert.Greater(t, resp.Count, 0, "Should find Luke")
				// First result should contain "Luke"
				if len(resp.Results) > 0 {
					assert.Contains(t, resp.Results[0].Name, "Luke")
				}
			},
		},
		{
			name:           "sort people by name",
			url:            "/people?sortBy=name&sortOrder=asc",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var resp domain.PaginatedResponse[domain.Person]
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				require.NoError(t, err)

				assert.NotEmpty(t, resp.Results)
				// Verify sorting
				if len(resp.Results) > 1 {
					assert.LessOrEqual(t, resp.Results[0].Name, resp.Results[1].Name,
						"Results should be sorted by name")
				}
			},
		},
		{
			name:           "pagination page 2",
			url:            "/people?page=2",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var resp domain.PaginatedResponse[domain.Person]
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				require.NoError(t, err)

				assert.Equal(t, 2, resp.Page)
				assert.NotEmpty(t, resp.Results, "Page 2 should have results")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create router with middleware for each test
			router := gin.New()
			router.Use(middleware.PaginationMiddleware())
			router.Use(middleware.QueryMiddleware())
			router.GET("/people", handler.ListPeople)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tt.url, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.checkResponse != nil {
				tt.checkResponse(t, w)
			}
		})
	}
}
