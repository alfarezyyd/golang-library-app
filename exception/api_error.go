package exception

type APIError struct {
	Status  int    `json:"status"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

func (apiError *APIError) Error() string {
	return apiError.Details
}

func NewAPIError(statusError int, titleError string, detailError string) *APIError {
	return &APIError{Status: statusError, Title: titleError, Details: detailError}
}
