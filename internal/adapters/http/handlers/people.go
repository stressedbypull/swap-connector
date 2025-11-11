package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/adapters/http/response"
	"github.com/stressedbypull/swapi-connector/internal/ports"
)

// PeopleHandler handles HTTP requests for people resources.
type PeopleHandler struct {
	service ports.PeopleServiceInterface
}

// NewPeopleHandler creates a new people handler with dependency injection.
func NewPeopleHandler(service ports.PeopleServiceInterface) *PeopleHandler {
	return &PeopleHandler{
		service: service,
	}
}

// ListPeople handles GET /people?page=1&search=luke&sortBy=name&sortOrder=asc
func (h *PeopleHandler) ListPeople(c *gin.Context) {
	// Parse and validate query parameters
	params := ParsePeopleQueryParams(c)

	// Call service
	result, err := h.service.ListPeople(
		c.Request.Context(),
		params.Page,
		params.Search,
		params.SortBy,
		params.SortOrder,
	)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, result)
}
