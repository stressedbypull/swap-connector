package errors

// APIError represents a domain error with HTTP status code.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"` // Don't expose HTTP status in JSON
}

// Error implements the error interface.
func (e APIError) Error() string {
	return e.Message
}

// Domain errors for SWAPI connector
var (
	// ErrPersonNotFound indicates a person was not found in SWAPI
	ErrPersonNotFound = APIError{
		Code:    "PERSON_NOT_FOUND",
		Message: "Person not found",
		Status:  404,
	}

	// ErrPlanetNotFound indicates a planet was not found in SWAPI
	ErrPlanetNotFound = APIError{
		Code:    "PLANET_NOT_FOUND",
		Message: "Planet not found",
		Status:  404,
	}

	// ErrInvalidSortField indicates an invalid sort field was provided
	ErrInvalidSortField = APIError{
		Code:    "INVALID_SORT_FIELD",
		Message: "Invalid sort field",
		Status:  400,
	}

	// ErrInvalidSortOrder indicates an invalid sort order was provided
	ErrInvalidSortOrder = APIError{
		Code:    "INVALID_SORT_ORDER",
		Message: "Sort order must be 'asc' or 'desc'",
		Status:  400,
	}

	// ErrInvalidPage indicates an invalid page number was provided
	ErrInvalidPage = APIError{
		Code:    "INVALID_PAGE",
		Message: "Page must be a positive integer",
		Status:  400,
	}

	// ErrSWAPIUnavailable indicates SWAPI service is unavailable
	ErrSWAPIUnavailable = APIError{
		Code:    "SWAPI_UNAVAILABLE",
		Message: "SWAPI service is temporarily unavailable",
		Status:  503,
	}

	// ErrRateLimitExceeded indicates rate limit was exceeded
	ErrRateLimitExceeded = APIError{
		Code:    "RATE_LIMIT_EXCEEDED",
		Message: "Rate limit exceeded, please try again later",
		Status:  429,
	}
)
