package types

// ErrorResponse is the response body of API errors
// swagger:model
type ErrorResponse struct {
	// The error message
	// required: true
	Message string `json:"message"`
}

// BadParameterError -
// swagger:model
type BadParameterError ErrorResponse

// InternalServerError -
// swagger:model
type InternalServerError ErrorResponse
