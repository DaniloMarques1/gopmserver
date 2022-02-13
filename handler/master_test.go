package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danilomarques1/gopmserver/dto"
	"github.com/danilomarques1/gopmserver/util"
	"github.com/go-chi/chi/v5"
)

type MasterServiceMock struct {
}

func (msm *MasterServiceMock) Save(masterDto *dto.MasterRequestDto) error {
	if masterDto.Email == "fitz@mail.com" {
		return util.NewApiError("E-mail already in use", http.StatusBadRequest)
	}
	return nil
}

func (msm *MasterServiceMock) Session(masterDto *dto.SessionRequestDto) (*dto.SessionResponseDto, error) {
	if masterDto.Pwd != "123456" {
		return nil, util.NewApiError("Invalid password", http.StatusUnauthorized)
	}
	response := dto.SessionResponseDto{Token: "token"}
	return &response, nil
}

func TestSaveMaster(t *testing.T) {
	cases := []struct {
		label           string
		body            string
		expectedStatus  int
		expectedMessage string
	}{
		{"TestSaveMaster", `{"email": "test@mail.com", "pwd": "secret"}`, http.StatusCreated, ""},
		{"TestSaveMasterErrEmailAlreadyInUse", `{"email": "fitz@mail.com", "pwd": "secret"}`, http.StatusBadRequest, "E-mail already in use"},
		{"TestSaveMasterErrInvalidBody", `invalidbody`, http.StatusBadRequest, "Invalid body"},
	}

	router := chi.NewRouter()
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			masterHandler := NewMasterHandler(&MasterServiceMock{})
			request := httptest.NewRequest(http.MethodPost, "/master", strings.NewReader(tc.body))

			rr := httptest.NewRecorder()
			router.Post("/master", masterHandler.Save)
			router.ServeHTTP(rr, request)

			if rr.Code != tc.expectedStatus {
				t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", tc.expectedStatus, rr.Code)
			}

			if tc.expectedMessage != "" {
				var response dto.ErrorDto
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Expect error to be nil when parsing response: %v\n", err)
				}

				if response.Message != tc.expectedMessage {
					t.Fatalf("Wrong message returned. Expected: %v got: %v\n", tc.expectedMessage, response.Message)
				}
			}

		})
	}
}

func TestSessionMaster(t *testing.T) {
	cases := []struct {
		label           string
		body            string
		expectedStatus  int
		expectedMessage string
	}{
		{"TestSessionMaster", `{"email": "test@mail.com", "pwd": "123456"}`, http.StatusOK, "token"},
		{"TestSessionMasterErrInvalidPwd", `{"email": "test@mail.com", "pwd": "invalid"}`, http.StatusUnauthorized, "Invalid password"},
		{"TestSessionMasterErrInvalidBody", `invalidbody`, http.StatusBadRequest, "Invalid body"},
	}

	router := chi.NewRouter()
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			masterHandler := NewMasterHandler(&MasterServiceMock{})
			request := httptest.NewRequest(http.MethodPost, "/session", strings.NewReader(tc.body))

			rr := httptest.NewRecorder()
			router.Post("/session", masterHandler.Session)
			router.ServeHTTP(rr, request)

			if rr.Code != tc.expectedStatus {
				t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", tc.expectedStatus, rr.Code)
			}

			if rr.Code == http.StatusOK {
				var response dto.SessionResponseDto
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Expect error to be nil when parsing response: %v\n", err)
				}

				if response.Token != tc.expectedMessage {
					t.Fatalf("Wrong message returned. Expected: %v got: %v\n", tc.expectedMessage, response.Token)
				}

			} else {
				var response dto.ErrorDto
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Expect error to be nil when parsing response: %v\n", err)
				}

				if response.Message != tc.expectedMessage {
					t.Fatalf("Wrong message returned. Expected: %v got: %v\n", tc.expectedMessage, response.Message)
				}
			}

		})
	}
}
