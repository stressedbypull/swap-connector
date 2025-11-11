package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	apierrors "github.com/stressedbypull/swapi-connector/internal/errors"
)

// HTTP status code constants
const (
	StatusOK                  = http.StatusOK
	StatusCreated             = http.StatusCreated
	StatusBadRequest          = http.StatusBadRequest
	StatusNotFound            = http.StatusNotFound
	StatusInternalServerError = http.StatusInternalServerError
)

// SuccessResponse represents a successful API response.
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// ErrorResponse represents an error API response.
// @name ErrorResponse
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains error information.
// @name ErrorDetail
type ErrorDetail struct {
	Message string                 `json:"message" example:"Person not found"`
	Code    string                 `json:"code,omitempty" example:"NOT_FOUND"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// OK sends a successful response with data.
func OK(c *gin.Context, data interface{}) {
	c.JSON(StatusOK, data)
}

// Created sends a 201 Created response.
func Created(c *gin.Context, data interface{}) {
	c.JSON(StatusCreated, SuccessResponse{Data: data})
}

// BadRequest sends a 400 Bad Request error.
func BadRequest(c *gin.Context, message string) {
	c.JSON(StatusBadRequest, ErrorResponse{
		Error: ErrorDetail{
			Message: message,
			Code:    "BAD_REQUEST",
		},
	})
}

// NotFound sends a 404 Not Found error.
func NotFound(c *gin.Context, message string) {
	c.JSON(StatusNotFound, ErrorResponse{
		Error: ErrorDetail{
			Message: message,
			Code:    "NOT_FOUND",
		},
	})
}

// InternalError sends a 500 Internal Server Error.
func InternalError(c *gin.Context, message string) {
	c.JSON(StatusInternalServerError, ErrorResponse{
		Error: ErrorDetail{
			Message: message,
			Code:    "INTERNAL_ERROR",
		},
	})
}

// ValidationError sends a 400 with validation details.
func ValidationError(c *gin.Context, details map[string]interface{}) {
	c.JSON(StatusBadRequest, ErrorResponse{
		Error: ErrorDetail{
			Message: "Validation failed",
			Code:    "VALIDATION_ERROR",
			Details: details,
		},
	})
}

// HandleError handles domain errors and sends appropriate HTTP responses.
// It checks if the error is an APIError and uses its status code and details,
// otherwise defaults to 500 Internal Server Error.
func HandleError(c *gin.Context, err error) {
	// Check if error is an APIError
	var apiErr apierrors.APIError
	if errors.As(err, &apiErr) {
		c.JSON(apiErr.Status, ErrorResponse{
			Error: ErrorDetail{
				Message: apiErr.Message,
				Code:    apiErr.Code,
			},
		})
		return
	}

	// Default to internal server error for unknown errors
	InternalError(c, err.Error())
}
