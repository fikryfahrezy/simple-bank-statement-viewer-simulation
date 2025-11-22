package http_server

import (
	"net/http"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/app_error"
)

// APIResponse represents a standard API response
type APIResponse struct {
	Message     string         `json:"message"`
	Error       string         `json:"error"`
	ErrorFields map[string]any `json:"error_fields"`
	Result      any            `json:"result"`
}

func SuccessResponse(w http.ResponseWriter, message string, data any) {
	JSON(w, http.StatusOK, APIResponse{
		Message: message,
		Error:   "",
		Result:  data,
	})
}

func CreatedResponse(w http.ResponseWriter, message string, data any) {
	JSON(w, http.StatusCreated, APIResponse{
		Message: message,
		Error:   "",
		Result:  data,
	})
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	errorCode := http.StatusText(statusCode)
	errorMessage := message
	var errorFields map[string]any

	if err != nil {
		// Extract error code if it's an AppError
		code := app_error.GetCode(err)
		if code != "" {
			errorCode = code
		}

		// Extract error message if it's an AppError
		message := app_error.GetMessage(err)
		if message != "" {
			errorMessage = message
		}

		// You can add error fields parsing here if needed
		errorFields = map[string]any{}
	}

	JSON(w, statusCode, APIResponse{
		Message:     errorMessage,
		Error:       errorCode,
		ErrorFields: errorFields,
	})
}

func BadRequestResponse(w http.ResponseWriter, message string, err error) {
	ErrorResponse(w, http.StatusBadRequest, message, err)
}

func InternalServerErrorResponse(w http.ResponseWriter, message string, err error) {
	ErrorResponse(w, http.StatusInternalServerError, message, err)
}

// ValidationErrorResponse creates a response for validation errors with field-specific errors
func ValidationErrorResponse(w http.ResponseWriter, message string, errorFields map[string]any) {
	JSON(w, http.StatusUnprocessableEntity, APIResponse{
		Message:     message,
		Error:       http.StatusText(http.StatusUnprocessableEntity),
		ErrorFields: errorFields,
	})
}
