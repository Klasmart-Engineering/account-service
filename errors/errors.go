package api_errors

import (
	"fmt"
)

// External API error struct, used for APIErrors returned to the client
type APIErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     string `json:"error"`
}

// Internal API error struct
type APIError struct {
	Status  int
	Code    string
	Message string
	Err     error
}

func (r *APIError) Error() string {
	return fmt.Sprintf("code %s: message %s error: %s", r.Code, r.Message, r.Err.Error())
}

func (e *APIError) ErrorResponse() *APIErrorResponse {
	response := &APIErrorResponse{
		Status:  e.Status,
		Code:    e.Code,
		Message: e.Message,
	}

	if e.Err != nil {
		response.Err = e.Err.Error()
	}
	return response
}

const (
	ErrCodeNotFound   = "NOT_FOUND"
	ErrCodeBadRequest = "BAD_REQUEST"
	ErrCodeUnknown    = "UNKNOWN"
	ErrCodeDbError    = "DATABASE_ERROR"
)

const (
	ErrMsgNotFound   = "%s with ID %s not found"
	ErrMsgBadRequest = "invalid request"
	ErrMsgUnknown    = "an unknown error occurred"
	ErrMsgDbError    = "an unknown database error occurred"
)
