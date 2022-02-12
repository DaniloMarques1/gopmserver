package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/danilomarques1/gopmserver/dto"
)

type ResponseWriterMock struct {
	statusCode int
	b          []byte
}

func NewResponseWriterMock() *ResponseWriterMock {
	rwm := &ResponseWriterMock{}
	return rwm
}

func (rwm *ResponseWriterMock) WriteHeader(statusCode int) {
	rwm.statusCode = statusCode
}

func (rwm *ResponseWriterMock) Write(src []byte) (int, error) {
	rwm.b = make([]byte, len(src))
	n := copy(rwm.b, src)
	return n, nil
}

func (rwm *ResponseWriterMock) Read() []byte {
	return rwm.b
}

func (rwm ResponseWriterMock) Header() http.Header {
	header := http.Header{}
	return header
}

func TestRespondJSON(t *testing.T) {
	rwm := NewResponseWriterMock()
	msg := "User stored"
	code := http.StatusCreated
	RespondJSON(rwm, msg, code)

	if rwm.statusCode != code {
		t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", code, rwm.statusCode)
	}

	var response dto.ErrorDto
	if err := json.Unmarshal(rwm.b, &response); err != nil {
		t.Fatalf("Error parsing the response that written: %v\n", err)
	}

	if response.Message != msg {
		t.Fatalf("Wrong message returned. Expect: %v got: %v\n", msg, response.Message)
	}
}

func TestRespondErrBadRequest(t *testing.T) {
	errMsg := "Invalid body"
	err := NewApiError(errMsg, http.StatusBadRequest)
	rwm := NewResponseWriterMock()
	RespondERR(rwm, err)
	if rwm.statusCode != http.StatusBadRequest {
		t.Fatalf("Wrong status returned. \n\tExpect: %v \n\tgot: %v\n", http.StatusBadRequest, rwm.statusCode)
	}

	var response dto.ErrorDto
	if err := json.Unmarshal(rwm.b, &response); err != nil {
		t.Fatalf("Err should be nil got: %v", err)
	}

	if response.Message != errMsg {
		t.Fatalf("Wrong msg returned. \n\tExpect: %v \n\tgot: %v\n", errMsg, response.Message)
	}
}

func TestRespondErrInternalServerError(t *testing.T) {
	errMsg := "Some crazy error"
	err := errors.New(errMsg)
	rwm := NewResponseWriterMock()
	RespondERR(rwm, err)
	if rwm.statusCode != http.StatusInternalServerError {
		t.Fatalf("Wrong status returned. \n\tExpect: %v \n\tgot: %v\n", http.StatusInternalServerError, rwm.statusCode)
	}

	var response dto.ErrorDto
	if err := json.Unmarshal(rwm.b, &response); err != nil {
		t.Fatalf("Err should be nil got: %v", err)
	}

	if response.Message != errMsg {
		t.Fatalf("Wrong msg returned. \n\tExpect: %v \n\tgot: %v\n", errMsg, response.Message)
	}
}
