package errors

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
