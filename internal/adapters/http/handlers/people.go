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

// ListPeople godoc
// @Summary      List Star Wars people
// @Description  Get a paginated list of people from SWAPI with optional search and sorting
// @Tags         people
// @Accept       json
// @Produce      json
// @Param        page       query     int     false  "Page number"           default(1)       example(1)
// @Param        search     query     string  false  "Search by name"        example(luke)
// @Param        sortBy     query     string  false  "Sort field"            Enums(name, created, mass)  example(name)
// @Param        sortOrder  query     string  false  "Sort order"            Enums(asc, desc)            default(asc)  example(asc)
// @Success      200  {object}  PeopleListResponse  "Successful response with people list"
// @Failure      400  {object}  ErrorResponse       "Invalid request parameters"
// @Failure      404  {object}  ErrorResponse       "Person not found"
// @Failure      500  {object}  ErrorResponse       "Internal server error"
// @Router       /people [get]
func (h *PeopleHandler) ListPeople(c *gin.Context) {
	// Parse and validate query parameters
	params, ok := ParsePeopleQueryParams(c)
	if !ok {
		return // Validation error already sent
	}

	// Call service
	result, err := h.service.ListPeople(
		c.Request.Context(),
		params.Page,
		params.Search,
		params.SortBy,
		params.SortOrder,
	)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.OK(c, result)
}
