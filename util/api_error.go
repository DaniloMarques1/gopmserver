package util

type ApiError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewApiError(msg string, status int) *ApiError {
	return &ApiError{Message: msg, StatusCode: status}
}

func (ae *ApiError) Error() string {
	return ae.Message
}
