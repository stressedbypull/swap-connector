package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stressedbypull/swapi-connector/internal/ports"
)

type PeopleHandler struct {
	service ports.PeopleServiceInterface
}

func NewPeopleHandler(s ports.PeopleServiceInterface) *PeopleHandler {
	return &PeopleHandler{
		service: s,
	}
}

// ListPeople handles GET /people?page=1&search=luke&sortBy=name&sortOrder=asc
func (h *PeopleHandler) ListPeople(c *gin.Context) {
	// Parse page number (default to 1)
	page := 1
	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	// Get query parameters
	search := c.Query("search")
	sortBy := c.Query("sortBy")
	sortOrder := c.Query("sortOrder")

	// Call service
	result, err := h.service.ListPeople(c.Request.Context(), page, 15, search, sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
