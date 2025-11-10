package swapi

import (
	"net/http"
	"time"
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
