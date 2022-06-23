package api_errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	pq "github.com/lib/pq"
)

// go-gin error handling middleware
func ErrorHandler(c *gin.Context) {
	c.Next()

	errors := []APIErrorResponse{}

	for _, err := range c.Errors {
		var errorResponse APIErrorResponse
		switch v := err.Err.(type) {
		// known API errors
		case *APIError:
			errorResponse = *v.ErrorResponse()
		// go-gin validation errors
		case validator.ValidationErrors:
			errorResponse = APIErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    ErrCodeBadRequest,
				Message: ErrMsgBadRequest,
				Err:     err.Error(),
			}
		// postgres errors
		case *pq.Error:
			errorResponse = APIErrorResponse{
				Status:  http.StatusInternalServerError,
				Code:    ErrCodeDbError,
				Message: ErrMsgDbError,
				Err:     err.Err.Error(),
			}
		// default to HTTP 500 unknown error
		default:
			errorResponse = APIErrorResponse{
				Status:  http.StatusInternalServerError,
				Code:    ErrCodeUnknown,
				Message: ErrMsgUnknown,
				Err:     err.Err.Error(),
			}
		}
		errors = append(errors, errorResponse)
	}

	if len(errors) > 0 {
		// use the status code from the last error
		finalStatus := errors[len(errors)-1].Status
		c.JSON(finalStatus, gin.H{"errors": errors})
	}
}
