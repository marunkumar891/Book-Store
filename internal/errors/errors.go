package errors

import "net/http"

type RestAPIError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NO_ERROR() RestAPIError {
	var restErr RestAPIError
	return restErr
}
func HasError(restErr *RestAPIError) bool {
	return *restErr != NO_ERROR()
}
func NewInternalServerError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal server error",
	}
}
func NewBadRequestError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad request",
	}
}
func NewNotFoundError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not found",
	}
}
