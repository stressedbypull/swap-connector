package swapi

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

const (
	OurPageSize   = 15
	SWAPIPageSize = 10
)

// Client implements ports.PeopleRepository and ports.PlanetsRepository.
// Uses dependency injection: httpClient is injected (or created if nil).
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a SWAPI repository client with dependency injection.
// If httpClient is nil, a default one is created.
func NewClient(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}
	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

// APIRetrievePeople fetches people with pagination (15 items per page).
// Fetches 2 SWAPI pages concurrently, maps to domain, and slices correctly.
func (c *Client) APIRetrievePeople(ctx context.Context, page int, search string) (domain.PaginatedResponse[domain.Person], error) {
	// Calculate which SWAPI pages to fetch
	startPage := ((page-1)*OurPageSize)/SWAPIPageSize + 1

	// Fetch 2 pages concurrently
	allPeople, totalCount, err := c.fetchPeoplePagesConcurrently(ctx, startPage, search)
	if err != nil {
		return domain.PaginatedResponse[domain.Person]{}, err
	}

	// Slice to get exactly the items for this page
	results := sliceForPage(allPeople, page)

	return domain.PaginatedResponse[domain.Person]{
		Count:    totalCount,
		Page:     page,
		PageSize: len(results),
		Results:  results,
	}, nil
}

// APIRetrievePersonByID fetches a single person by ID.
func (c *Client) APIRetrievePersonByID(ctx context.Context, id string) (domain.Person, error) {
	// TODO: Implement
	return domain.Person{}, nil
}

// fetchPeoplePagesConcurrently fetches 2 SWAPI pages concurrently.
func (c *Client) fetchPeoplePagesConcurrently(ctx context.Context, startPage int, search string) ([]domain.Person, int, error) {
	type result struct {
		page   int
		people []domain.Person
		count  int
		err    error
	}

	results := make(chan result, 2)
	var wg sync.WaitGroup

	// Fetch 2 pages concurrently
	for i := 0; i < 2; i++ {
		wg.Add(1)
		pageNum := startPage + i

		go func(p int) {
			defer wg.Done()

			// Fetch raw DTO
			dto, err := c.fetchPeoplePageHTTP(ctx, p, search)
			if err != nil {
				results <- result{page: p, err: err}
				return
			}

			// Map to domain
			people := MapPeopleToDomain(dto.Results)

			results <- result{
				page:   p,
				people: people,
				count:  dto.Count,
			}
		}(pageNum)
	}

	// Close channel when done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results in order
	pageResults := make(map[int][]domain.Person)
	var totalCount int

	for res := range results {
		if res.err != nil {
			return nil, 0, res.err
		}
		pageResults[res.page] = res.people
		totalCount = res.count
	}

	// Combine in order
	var allPeople []domain.Person
	for i := 0; i < 2; i++ {
		if people, exists := pageResults[startPage+i]; exists {
			allPeople = append(allPeople, people...)
		}
	}

	return allPeople, totalCount, nil
}

// fetchPeoplePageHTTP performs the HTTP call to SWAPI (raw DTO only).
func (c *Client) fetchPeoplePageHTTP(ctx context.Context, page int, search string) (*SWAPIPeopleResponse, error) {
	url := BuildURL(c.baseURL, "people", page, search)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var swapiResp SWAPIPeopleResponse
	if err := json.NewDecoder(resp.Body).Decode(&swapiResp); err != nil {
		return nil, err
	}

	return &swapiResp, nil
}

// sliceForPage extracts the correct 15 items for the requested page.
func sliceForPage(items []domain.Person, page int) []domain.Person {
	if len(items) == 0 {
		return []domain.Person{}
	}

	// Calculate offset within fetched items (~20 items)
	offset := ((page - 1) * OurPageSize) % (SWAPIPageSize * 2)
	end := offset + OurPageSize

	if offset >= len(items) {
		return []domain.Person{}
	}
	if end > len(items) {
		end = len(items)
	}

	return items[offset:end]
}
