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

func TestRespondERR(t *testing.T) {
	cases := []struct {
		label      string
		err        error
		expectCode int
	}{
		{
			"TestRespondERRInternalServerError",
			errors.New("Something wen wrong"),
			http.StatusInternalServerError,
		},
		{
			"TestRespondERRBadRequest",
			NewApiError("Invalid body", http.StatusBadRequest),
			http.StatusBadRequest,
		},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			rwm := NewResponseWriterMock()
			RespondERR(rwm, tc.err)
			if rwm.statusCode != tc.expectCode {
				t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", tc.expectCode, rwm.statusCode)
			}
			var response dto.ErrorDto
			if err := json.Unmarshal(rwm.b, &response); err != nil {
				t.Fatalf("Error parsing the response that written: %v\n", err)
			}

			if response.Message != tc.err.Error() {
				t.Fatalf("Wrong message returned. Expect: %v got: %v\n", tc.err.Error(), response.Message)
			}
		})
	}
}
