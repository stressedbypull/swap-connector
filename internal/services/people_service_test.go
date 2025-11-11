package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	errDomain "github.com/stressedbypull/swapi-connector/internal/errors"
	"github.com/stressedbypull/swapi-connector/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPeopleService_ListPeople(t *testing.T) {
	// Setup test data
	date1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	date3 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name              string
		page              int
		searchTerm        string
		sortBy            string
		sortOrder         string
		mockResponse      domain.PaginatedResponse[domain.Person]
		mockError         error
		setupContext      func() context.Context
		wantError         bool
		wantErrorContains string
		wantCount         int
		validateResults   func(t *testing.T, results []domain.Person)
	}{
		{
			name:       "Success - Sort by name ascending",
			page:       1,
			searchTerm: "",
			sortBy:     "name",
			sortOrder:  "asc",
			mockResponse: domain.PaginatedResponse[domain.Person]{
				Count:    2,
				Page:     1,
				PageSize: 2,
				Results: []domain.Person{
					{Name: "Luke Skywalker", Mass: 77, Create: time.Now()},
					{Name: "Darth Vader", Mass: 136, Create: time.Now()},
				},
			},
			mockError:    nil,
			setupContext: func() context.Context { return context.Background() },
			wantError:    false,
			wantCount:    2,
			validateResults: func(t *testing.T, results []domain.Person) {
				assert.Equal(t, "Darth Vader", results[0].Name, "First should be Darth Vader")
				assert.Equal(t, "Luke Skywalker", results[1].Name, "Second should be Luke Skywalker")
			},
		},
		{
			name:       "Success - Search filter",
			page:       1,
			searchTerm: "sky",
			sortBy:     "",
			sortOrder:  "asc",
			mockResponse: domain.PaginatedResponse[domain.Person]{
				Count:    3,
				Page:     1,
				PageSize: 3,
				Results: []domain.Person{
					{Name: "Luke Skywalker", Mass: 77, Create: time.Now()},
					{Name: "Darth Vader", Mass: 136, Create: time.Now()},
					{Name: "Anakin Skywalker", Mass: 84, Create: time.Now()},
				},
			},
			mockError:    nil,
			setupContext: func() context.Context { return context.Background() },
			wantError:    false,
			wantCount:    2,
			validateResults: func(t *testing.T, results []domain.Person) {
				for _, person := range results {
					assert.Contains(t, person.Name, "Skywalker", "All results should contain 'Skywalker'")
				}
			},
		},
		{
			name:       "Success - Sort by mass descending",
			page:       1,
			searchTerm: "",
			sortBy:     "mass",
			sortOrder:  "desc",
			mockResponse: domain.PaginatedResponse[domain.Person]{
				Count:    3,
				Page:     1,
				PageSize: 3,
				Results: []domain.Person{
					{Name: "Luke Skywalker", Mass: 77, Create: time.Now()},
					{Name: "Darth Vader", Mass: 136, Create: time.Now()},
					{Name: "Leia Organa", Mass: 49, Create: time.Now()},
				},
			},
			mockError:    nil,
			setupContext: func() context.Context { return context.Background() },
			wantError:    false,
			wantCount:    3,
			validateResults: func(t *testing.T, results []domain.Person) {
				assert.Equal(t, 136, results[0].Mass, "First should be heaviest")
				assert.Equal(t, 77, results[1].Mass, "Second should be medium")
				assert.Equal(t, 49, results[2].Mass, "Third should be lightest")
			},
		},
		{
			name:       "Success - Sort by created date ascending",
			page:       1,
			searchTerm: "",
			sortBy:     "created",
			sortOrder:  "asc",
			mockResponse: domain.PaginatedResponse[domain.Person]{
				Count:    3,
				Page:     1,
				PageSize: 3,
				Results: []domain.Person{
					{Name: "Person2", Mass: 80, Create: date2},
					{Name: "Person3", Mass: 90, Create: date3},
					{Name: "Person1", Mass: 70, Create: date1},
				},
			},
			mockError:    nil,
			setupContext: func() context.Context { return context.Background() },
			wantError:    false,
			wantCount:    3,
			validateResults: func(t *testing.T, results []domain.Person) {
				assert.Equal(t, date1, results[0].Create, "First should be oldest")
				assert.Equal(t, date2, results[1].Create, "Second should be middle")
				assert.Equal(t, date3, results[2].Create, "Third should be newest")
			},
		},
		{
			name:              "Error - SWAPI Unavailable",
			page:              1,
			searchTerm:        "",
			sortBy:            "",
			sortOrder:         "asc",
			mockResponse:      domain.PaginatedResponse[domain.Person]{},
			mockError:         errors.New("connection refused"),
			setupContext:      func() context.Context { return context.Background() },
			wantError:         true,
			wantErrorContains: "connection refused",
			wantCount:         0,
			validateResults:   nil,
		},
		{
			name:              "Error - Person Not Found",
			page:              99,
			searchTerm:        "",
			sortBy:            "",
			sortOrder:         "asc",
			mockResponse:      domain.PaginatedResponse[domain.Person]{},
			mockError:         errDomain.ErrPersonNotFound,
			setupContext:      func() context.Context { return context.Background() },
			wantError:         true,
			wantErrorContains: "not found",
			wantCount:         0,
			validateResults:   nil,
		},
		{
			name:         "Error - Context Cancelled",
			page:         1,
			searchTerm:   "",
			sortBy:       "",
			sortOrder:    "asc",
			mockResponse: domain.PaginatedResponse[domain.Person]{},
			mockError:    context.Canceled,
			setupContext: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			},
			wantError:         true,
			wantErrorContains: "context canceled",
			wantCount:         0,
			validateResults:   nil,
		},
		{
			name:       "Edge Case - Empty Results",
			page:       1,
			searchTerm: "nonexistent",
			sortBy:     "",
			sortOrder:  "asc",
			mockResponse: domain.PaginatedResponse[domain.Person]{
				Count:    0,
				Page:     1,
				PageSize: 0,
				Results:  []domain.Person{},
			},
			mockError:       nil,
			setupContext:    func() context.Context { return context.Background() },
			wantError:       false,
			wantCount:       0,
			validateResults: nil,
		},
		{
			name:       "Edge Case - Invalid Sort Field (no sort applied)",
			page:       1,
			searchTerm: "",
			sortBy:     "invalid_field",
			sortOrder:  "asc",
			mockResponse: domain.PaginatedResponse[domain.Person]{
				Count:    2,
				Page:     1,
				PageSize: 2,
				Results: []domain.Person{
					{Name: "Zulu", Mass: 77, Create: time.Now()},
					{Name: "Alpha", Mass: 136, Create: time.Now()},
				},
			},
			mockError:    nil,
			setupContext: func() context.Context { return context.Background() },
			wantError:    false,
			wantCount:    2,
			validateResults: func(t *testing.T, results []domain.Person) {
				// Should maintain original order since invalid sort field
				assert.Equal(t, "Zulu", results[0].Name)
				assert.Equal(t, "Alpha", results[1].Name)
			},
		},
		{
			name:       "Edge Case - No Sort Field",
			page:       1,
			searchTerm: "",
			sortBy:     "",
			sortOrder:  "asc",
			mockResponse: domain.PaginatedResponse[domain.Person]{
				Count:    2,
				Page:     1,
				PageSize: 2,
				Results: []domain.Person{
					{Name: "Zulu", Mass: 77, Create: time.Now()},
					{Name: "Alpha", Mass: 136, Create: time.Now()},
				},
			},
			mockError:    nil,
			setupContext: func() context.Context { return context.Background() },
			wantError:    false,
			wantCount:    2,
			validateResults: func(t *testing.T, results []domain.Person) {
				// Should maintain original order
				assert.Equal(t, "Zulu", results[0].Name)
				assert.Equal(t, "Alpha", results[1].Name)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctx := tt.setupContext()
			mockRepo := mocks.NewMockSwapiRepository()
			mockRepo.On("APIRetrievePeople", mock.Anything, tt.page, tt.searchTerm).
				Return(tt.mockResponse, tt.mockError)
			service := NewPeopleService(mockRepo)

			// Act
			result, err := service.ListPeople(ctx, tt.page, tt.searchTerm, tt.sortBy, tt.sortOrder)

			// Assert
			if tt.wantError {
				assert.Error(t, err)
				if tt.wantErrorContains != "" {
					assert.Contains(t, err.Error(), tt.wantErrorContains)
				}
				assert.Empty(t, result.Results)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantCount, len(result.Results))
				if tt.validateResults != nil {
					tt.validateResults(t, result.Results)
				}
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestPeopleService_GetPeopleByID(t *testing.T) {
	tests := []struct {
		name              string
		personID          string
		mockPerson        domain.Person
		mockError         error
		wantError         bool
		wantErrorContains string
		validateResult    func(t *testing.T, person domain.Person)
	}{
		{
			name:     "Success - Get Person By ID",
			personID: "1",
			mockPerson: domain.Person{
				Name:   "Luke Skywalker",
				Mass:   77,
				Create: time.Now(),
				Films:  []string{"film1", "film2"},
			},
			mockError: nil,
			wantError: false,
			validateResult: func(t *testing.T, person domain.Person) {
				assert.Equal(t, "Luke Skywalker", person.Name)
				assert.Equal(t, 77, person.Mass)
				assert.Equal(t, 2, len(person.Films))
			},
		},
		{
			name:              "Error - Person Not Found",
			personID:          "9999",
			mockPerson:        domain.Person{},
			mockError:         errDomain.ErrPersonNotFound,
			wantError:         true,
			wantErrorContains: "not found",
			validateResult:    nil,
		},
		{
			name:              "Error - SWAPI Unavailable",
			personID:          "1",
			mockPerson:        domain.Person{},
			mockError:         errDomain.ErrSWAPIUnavailable,
			wantError:         true,
			wantErrorContains: "unavailable",
			validateResult:    nil,
		},
		{
			name:              "Error - Generic Error",
			personID:          "1",
			mockPerson:        domain.Person{},
			mockError:         errors.New("network timeout"),
			wantError:         true,
			wantErrorContains: "timeout",
			validateResult:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctx := context.Background()
			mockRepo := mocks.NewMockSwapiRepository()
			mockRepo.On("APIRetrievePersonByID", ctx, tt.personID).
				Return(tt.mockPerson, tt.mockError)
			service := NewPeopleService(mockRepo)

			// Act
			result, err := service.GetPeopleByID(ctx, tt.personID)

			// Assert
			if tt.wantError {
				assert.Error(t, err)
				if tt.wantErrorContains != "" {
					assert.Contains(t, err.Error(), tt.wantErrorContains)
				}
				assert.Empty(t, result.Name)
			} else {
				assert.NoError(t, err)
				if tt.validateResult != nil {
					tt.validateResult(t, result)
				}
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
