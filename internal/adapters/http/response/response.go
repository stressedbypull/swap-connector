package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains error information.
type ErrorDetail struct {
	Message string                 `json:"message"`
	Code    string                 `json:"code,omitempty"`
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
