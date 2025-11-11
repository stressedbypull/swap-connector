package swapi

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/pagination"
)

const (
	defaultTimeout  = 15 * time.Second
	defaultPageSize = 15
)

// Client implements ports.PeopleRepository and ports.PlanetsRepository.
// It fetches data from SWAPI, maps DTOs to domain objects.
type Client struct {
	baseURL    string
	httpClient *http.Client
	pageSize   int // Desired page size for responses (SWAPI returns ~10 per page)
}

// NewClient creates a SWAPI client with dependency injection.
// If httpClient is nil, a default client with 15s timeout is created.
// Page size is read from SWAPI_PAGE_SIZE environment variable (default: 15).
func NewClient(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: defaultTimeout,
		}
	}

	// Read page size from environment variable
	pageSize := defaultPageSize
	if envPageSize := os.Getenv("SWAPI_PAGE_SIZE"); envPageSize != "" {
		if size, err := strconv.Atoi(envPageSize); err == nil && size > 0 {
			pageSize = size
		}
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
		pageSize:   pageSize,
	}
}

// APIRetrievePeople fetches people with pagination from SWAPI.
// Aggregates SWAPI pages (~10 items each) to return the configured page size.
func (c *Client) APIRetrievePeople(ctx context.Context, page int, search string) (domain.PaginatedResponse[domain.Person], error) {
	const swapiPageSize = 10 // SWAPI returns ~10 items per page

	// Fetch aggregated data from SWAPI
	allPeople, totalCount, err := c.fetchAggregatedPeople(ctx, page, search, swapiPageSize)
	if err != nil {
		return domain.PaginatedResponse[domain.Person]{}, err
	}

	// Use pagination package to build the response
	strategy := pagination.NewAggregationStrategy(page, c.pageSize, swapiPageSize)
	return pagination.BuildResponse(allPeople, totalCount, strategy), nil
}

// APIRetrievePersonByID fetches a single person by ID from SWAPI.
func (c *Client) APIRetrievePersonByID(ctx context.Context, id string) (domain.Person, error) {
	url := c.baseURL + "/people/" + id + "/"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.Person{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return domain.Person{}, err
	}
	defer resp.Body.Close()

	// Validate HTTP status code
	if resp.StatusCode != http.StatusOK {
		return domain.Person{}, handleHTTPError(resp.StatusCode, "person")
	}

	var personDTO PersonDTO
	if err := json.NewDecoder(resp.Body).Decode(&personDTO); err != nil {
		return domain.Person{}, err
	}

	person := MapPersonDTOToDomain(personDTO)
	return person, nil
}

// fetchAggregatedPeople fetches multiple SWAPI pages and aggregates them.
// This helper extracts the data-fetching logic from pagination logic.
func (c *Client) fetchAggregatedPeople(ctx context.Context, page int, search string, swapiPageSize int) ([]domain.Person, int, error) {
	// Create pagination strategy to determine which pages to fetch
	strategy := pagination.NewAggregationStrategy(page, c.pageSize, swapiPageSize)
	startPage, _, pagesNeeded := strategy.CalculatePageRange()

	var allPeople []domain.Person
	var totalCount int

	// Fetch all necessary SWAPI pages
	for i := 0; i < pagesNeeded; i++ {
		currentPage := startPage + i

		dto, err := c.fetchPeoplePage(ctx, currentPage, search)
		if err != nil {
			return nil, 0, err
		}

		// Store total count from first response
		if i == 0 {
			totalCount = dto.Count
		}

		// Map DTOs to domain objects
		people := MapPeopleToDomain(dto.Results)
		allPeople = append(allPeople, people...)

		// Stop if we got all available results (partial page)
		if len(dto.Results) < swapiPageSize {
			break
		}
	}

	return allPeople, totalCount, nil
}

// fetchPeoplePage performs HTTP request to SWAPI people endpoint.
func (c *Client) fetchPeoplePage(ctx context.Context, page int, search string) (*SWAPIPeopleResponse, error) {
	url := BuildURL(c.baseURL, "people", page, search)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Validate HTTP status code
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			// For people list endpoint, 404 might mean empty results
			return &SWAPIPeopleResponse{Count: 0, Results: []PersonDTO{}}, nil
		}
		return nil, handleHTTPErrorForList(resp.StatusCode)
	}

	var response SWAPIPeopleResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
