package swapi

import (
	"context"
	"net/http"
	"time"

	"github.com/stressedbypull/swapi-connector/internal/domain"
)

// Client implements ports.PeopleRepository and ports.PlanetsRepository
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient returns a new SWAPI client. The http.Client is reused and safe for concurrent use.
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *Client) FetchPeople(ctx context.Context, page, pageSize int, search string) (domain.PaginatedResponse[domain.Person], error) {

	return domain.PaginatedResponse[domain.Person]{}, nil
}
