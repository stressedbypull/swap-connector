package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/middleware"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/response"
	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/mocks"
	"github.com/stressedbypull/swapi-connector/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TestPeopleHandler_ListPeople - Unit test with mocks
func TestPeopleHandler_ListPeople(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		url            string
		setupMock      func(m *mocks.MockSwapiRepository)
		expectedStatus int
		checkResponse  func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name: "list people successfully",
			url:  "/people",
			setupMock: func(m *mocks.MockSwapiRepository) {
				mockResp := domain.PaginatedResponse[domain.Person]{
					Count: 2,
					Page:  1,
					Results: []domain.Person{
						{Name: "Luke Skywalker", Mass: 77, Films: []string{"film1"}},
						{Name: "Darth Vader", Mass: 136, Films: []string{"film2"}},
					},
				}
				m.On("APIRetrievePeople", mock.Anything, 1, "").Return(mockResp, nil)
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var resp domain.PaginatedResponse[domain.Person]
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				require.NoError(t, err)

				assert.Equal(t, 2, resp.Count)
				assert.Equal(t, 1, resp.Page)
				assert.Len(t, resp.Results, 2)
				assert.Equal(t, "Luke Skywalker", resp.Results[0].Name)
			},
		},
		{
			name: "sort by name ascending",
			url:  "/people?sortBy=name&sortOrder=asc",
			setupMock: func(m *mocks.MockSwapiRepository) {
				mockResp := domain.PaginatedResponse[domain.Person]{
					Count: 2,
					Page:  1,
					Results: []domain.Person{
						{Name: "Leia Organa", Mass: 49, Films: []string{"film1"}},
						{Name: "Darth Vader", Mass: 136, Films: []string{"film2"}},
					},
				}
				m.On("APIRetrievePeople", mock.Anything, 1, "").Return(mockResp, nil)
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var resp domain.PaginatedResponse[domain.Person]
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				require.NoError(t, err)

				assert.Len(t, resp.Results, 2)
				// Should be sorted: Darth Vader, Leia Organa
				assert.Equal(t, "Darth Vader", resp.Results[0].Name)
				assert.Equal(t, "Leia Organa", resp.Results[1].Name)
			},
		},
		{
			name: "invalid sort field",
			url:  "/people?sortBy=invalid",
			setupMock: func(m *mocks.MockSwapiRepository) {
				// No mock setup needed - validation happens before repository call
			},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var resp response.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				require.NoError(t, err)

				assert.Equal(t, "VALIDATION_ERROR", resp.Error.Code)
				assert.Contains(t, resp.Error.Message, "Validation failed")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock for each test
			mockRepo := mocks.NewMockSwapiRepository()
			if tt.setupMock != nil {
				tt.setupMock(mockRepo)
			}

			service := services.NewPeopleService(mockRepo)
			handler := NewPeopleHandler(service)

			// Create router with middleware
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

			mockRepo.AssertExpectations(t)
		})
	}
}
