package swapi

import (
	"fmt"
	"net/http"

	"github.com/stressedbypull/swapi-connector/internal/errors"
)

// handleHTTPError maps HTTP status codes to domain errors.
// This centralizes error handling for all SWAPI HTTP requests.
func handleHTTPError(statusCode int, resourceType string) error {
	switch statusCode {
	case http.StatusNotFound:
		// Return appropriate error based on resource type
		if resourceType == "person" {
			return errors.ErrPersonNotFound
		}
		return errors.ErrPlanetNotFound

	case http.StatusTooManyRequests:
		return errors.ErrRateLimitExceeded

	case http.StatusServiceUnavailable, http.StatusBadGateway, http.StatusGatewayTimeout:
		return errors.ErrSWAPIUnavailable

	default:
		return fmt.Errorf("unexpected status code: %d", statusCode)
	}
}

// handleHTTPErrorForList handles errors for list endpoints (people, planets).
// For list endpoints, 404 might mean empty results rather than an error.
func handleHTTPErrorForList(statusCode int) error {
	switch statusCode {
	case http.StatusTooManyRequests:
		return errors.ErrRateLimitExceeded

	case http.StatusServiceUnavailable, http.StatusBadGateway, http.StatusGatewayTimeout:
		return errors.ErrSWAPIUnavailable

	default:
		return fmt.Errorf("unexpected status code: %d", statusCode)
	}
}
