package swapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/stressedbypull/swapi-connector/internal/domain"
	"github.com/stressedbypull/swapi-connector/internal/errors"
)

const (
	defaultTimeout = 15 * time.Second
)

// Client implements ports.PeopleRepository and ports.PlanetsRepository.
// It fetches data from SWAPI, maps DTOs to domain objects.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a SWAPI client with dependency injection.
// If httpClient is nil, a default client with 15s timeout is created.
func NewClient(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: defaultTimeout,
		}
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

// APIRetrievePeople fetches people with pagination from SWAPI.
// Returns SWAPI's pagination (~10 items per page).
func (c *Client) APIRetrievePeople(ctx context.Context, page int, search string) (domain.PaginatedResponse[domain.Person], error) {
	// Fetch from SWAPI
	dto, err := c.fetchPeoplePage(ctx, page, search)
	if err != nil {
		return domain.PaginatedResponse[domain.Person]{}, err
	}

	// Map DTOs to domain objects
	people := MapPeopleToDomain(dto.Results)

	response := domain.PaginatedResponse[domain.Person]{
		Count:    dto.Count,
		Page:     page,
		PageSize: len(people),
		Results:  people,
	}

	return response, nil
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
		switch resp.StatusCode {
		case http.StatusNotFound:
			return domain.Person{}, errors.ErrPersonNotFound
		case http.StatusTooManyRequests:
			return domain.Person{}, errors.ErrRateLimitExceeded
		case http.StatusServiceUnavailable, http.StatusBadGateway, http.StatusGatewayTimeout:
			return domain.Person{}, errors.ErrSWAPIUnavailable
		default:
			return domain.Person{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	}

	var personDTO PersonDTO
	if err := json.NewDecoder(resp.Body).Decode(&personDTO); err != nil {
		return domain.Person{}, err
	}

	person := MapPersonDTOToDomain(personDTO)
	return person, nil
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
		switch resp.StatusCode {
		case http.StatusNotFound:
			// For people list endpoint, 404 might mean empty results
			return &SWAPIPeopleResponse{Count: 0, Results: []PersonDTO{}}, nil
		case http.StatusTooManyRequests:
			return nil, errors.ErrRateLimitExceeded
		case http.StatusServiceUnavailable, http.StatusBadGateway, http.StatusGatewayTimeout:
			return nil, errors.ErrSWAPIUnavailable
		default:
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	}

	var response SWAPIPeopleResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
