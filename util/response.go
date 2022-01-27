package util

import (
	"encoding/json"
	"net/http"

	"github.com/danilomarques1/gopmserver/dto"
)

func RespondJSON(w http.ResponseWriter, message string, statusCode int) {
	errDto := dto.ErrorDto{Message: message}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errDto)
}

func RespondERR(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *ApiError:
		errApi := err.(*ApiError)
		RespondJSON(w, errApi.Message, errApi.StatusCode)
	default:
		RespondJSON(w, err.Error(), http.StatusInternalServerError)
	}
}
