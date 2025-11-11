package validation

import (
	"fmt"
	"strconv"
	"strings"
)

// Error represents a single validation error.
type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value,omitempty"`
}

// Validator collects validation errors.
type Validator struct {
	errors []Error
}

// New creates a new Validator.
func New() *Validator {
	return &Validator{
		errors: make([]Error, 0),
	}
}

// AddError adds a validation error.
func (v *Validator) AddError(field, message, value string) {
	v.errors = append(v.errors, Error{
		Field:   field,
		Message: message,
		Value:   value,
	})
}

// HasErrors returns true if there are validation errors.
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

// Errors returns all validation errors.
func (v *Validator) Errors() []Error {
	return v.errors
}

// ErrorsMap returns errors as a map for JSON response.
func (v *Validator) ErrorsMap() map[string]interface{} {
	errMap := make(map[string]interface{})
	for _, err := range v.errors {
		errMap[err.Field] = err.Message
	}
	return errMap
}

// ValidatePositiveInt validates that a string is a positive integer.
func (v *Validator) ValidatePositiveInt(field, value string) bool {
	num, err := strconv.Atoi(value)
	if err != nil || num <= 0 {
		v.AddError(field, "must be a positive integer", value)
		return false
	}
	return true
}

// ValidateNotEmpty validates that a string is not empty.
func (v *Validator) ValidateNotEmpty(field, value string) bool {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		v.AddError(field, "cannot be empty", value)
		return false
	}
	return true
}

// ValidateOneOf validates that value is one of the allowed values.
func (v *Validator) ValidateOneOf(field, value string, allowed []string) bool {
	if value == "" {
		return true // Optional field
	}

	for _, a := range allowed {
		if value == a {
			return true
		}
	}

	v.AddError(field, fmt.Sprintf("must be one of: %s", strings.Join(allowed, ", ")), value)
	return false
}

// ParseMass parses mass string to int, handling "unknown" and comma-separated values.
// Returns 0 for invalid values instead of error.
// Examples: "77" -> 77, "1,358" -> 1358, "unknown" -> 0
func ParseMass(s string) int {
	if s == "" || s == "unknown" {
		return 0
	}

	// Remove commas: "1,358" -> "1358"
	cleaned := strings.ReplaceAll(s, ",", "")

	// Parse as integer
	mass, err := strconv.Atoi(cleaned)
	if err != nil {
		// Try as float for values like "78.2"
		if massFloat, floatErr := strconv.ParseFloat(cleaned, 64); floatErr == nil {
			return int(massFloat)
		}
		return 0
	}

	return mass
}
