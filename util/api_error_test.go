package util

import (
	"net/http"
	"testing"
)

func TestNewApiError(t *testing.T) {
	errMsg := "Invalid body"
	err := NewApiError(errMsg, http.StatusBadRequest)

	if err == nil {
		t.Fatal("Should have returned an error\n")
	}
	if err.Error() != errMsg {
		t.Fatalf("Wrong message returned. Expected: %v got: %v\n", errMsg, err.Error())
	}
}
