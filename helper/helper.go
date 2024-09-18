package helper

import "github.com/go-playground/validator/v10"

// Response is a standardized format for API responses, containing meta information and data.
type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"` // Data holds the actual response data (could be of any type).
}

// Meta contains the metadata for an API response, including the HTTP status code, status, and message.
type Meta struct {
	Code    int    `json:"code"`    // HTTP status code
	Status  string `json:"status"`  // Status (e.g., "Success", "Failed")
	Message string `json:"message"` // Descriptive message
}

// APIResponse formats a consistent API response structure with a meta block and response data.
func APIResponse(code int, status string, message string, data any) Response {
	meta := Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}
	responseFormat := Response{
		Meta: meta,
		Data: data,
	}
	return responseFormat
}

// FormatValidationError processes validation errors and returns them as a slice of strings.
func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		// Convert each validation error to a human-readable string and append it to the errors slice.
		errors = append(errors, e.Error())
	}
	return errors
}
