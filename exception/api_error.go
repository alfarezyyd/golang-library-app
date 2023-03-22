package exception

type APIError struct {
	Status  int    `json:"status"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

func (apiError *APIError) Error() string {
	return apiError.Details
}

func NewAPIError(status int, title string, details string) *APIError {
	return &APIError{Status: status, Title: title, Details: details}
}
