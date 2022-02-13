package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danilomarques1/gopmserver/dto"
	"github.com/danilomarques1/gopmserver/util"
	"github.com/go-chi/chi/v5"
)

type PasswordServiceMock struct {
}

func (psm *PasswordServiceMock) Save(masterId string, pwdDto *dto.PasswordRequestDto) error {
	if pwdDto.Key == "facebook" {
		return util.NewApiError("Key already in use", http.StatusBadRequest)
	}
	return nil
}

func (psm *PasswordServiceMock) FindByKey(masterId, key string) (*dto.PasswordResponseDto, error) {
	if key == "github" {
		return nil, util.NewApiError("No password found with the given key", http.StatusNotFound)
	}

	return &dto.PasswordResponseDto{Id: "1", Key: key, Pwd: "123456"}, nil
}

func (psm *PasswordServiceMock) Keys(masterId string) (*dto.PasswordKeysDto, error) {
	if len(masterId) == 0 {
		return nil, errors.New("Some error")
	}
	keys := []string{"github", "orkut"}
	response := dto.PasswordKeysDto{Keys: keys}
	return &response, nil
}

func (psm *PasswordServiceMock) RemoveByKey(masterId, key string) error {
	if key == "github" {
		return util.NewApiError("No password found with the given key", http.StatusNotFound)
	}
	return nil
}

func (psm *PasswordServiceMock) UpdateByKey(masterId string, pwdDto *dto.PasswordUpdateRequestDto) error {
	if pwdDto.Key == "github" {
		return util.NewApiError("No password found with the given key", http.StatusNotFound)
	}
	return nil
}

const USERID = "1"

func TestPasswordSave(t *testing.T) {
	cases := []struct {
		label           string
		body            string
		expectedStatus  int
		expectedMessage string
	}{
		{"TestPasswordSave", `{"key": "orkut", "pwd": "123456"}`, http.StatusCreated, ""},
		{"TestPasswordSaveErrKeyAlreadyInUse", `{"key": "facebook", "pwd": "123456"}`, http.StatusBadRequest, "Key already in use"},
		{"TestPasswordSaveErrInvalidBody", `invalidbody`, http.StatusBadRequest, "Invalid body"},
	}

	router := chi.NewRouter()
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdHandler := NewPasswordHandler(&PasswordServiceMock{})
			request := httptest.NewRequest(http.MethodPost, "/password", strings.NewReader(tc.body))
			request.Header.Add("userId", USERID)

			rr := httptest.NewRecorder()
			router.Post("/password", pwdHandler.Save)
			router.ServeHTTP(rr, request)

			if rr.Code != tc.expectedStatus {
				t.Fatalf("Wrong status returned. Expect: %v got: %v\n", tc.expectedStatus, rr.Code)
			}

			if tc.expectedMessage != "" {
				var response dto.ErrorDto
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Expect err to be nil when parsing response. Instead got: %v\n", err)
				}
				if response.Message != tc.expectedMessage {
					t.Fatalf("Wrong message returned. Got: %v\n", response.Message)
				}
			}

		})
	}
}

func TestFindPassword(t *testing.T) {
	cases := []struct {
		label           string
		key             string
		expectedStatus  int
		expectedMessage string
	}{
		{"TestFindPassword", "orkut", http.StatusOK, "123456"},
		{"TestFindPasswordErrKeyNotFound", "github", http.StatusNotFound, "No password found with the given key"},
	}

	router := chi.NewRouter()
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdHandler := NewPasswordHandler(&PasswordServiceMock{})
			url := fmt.Sprintf("/password/%v", tc.key)
			request := httptest.NewRequest(http.MethodGet, url, nil)
			request.Header.Add("userId", USERID)

			rr := httptest.NewRecorder()
			router.Get("/password/{key}", pwdHandler.FindByKey)
			router.ServeHTTP(rr, request)

			if rr.Code != tc.expectedStatus {
				t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", tc.expectedStatus, rr.Code)
			}

			if rr.Code == http.StatusOK {
				var response dto.PasswordResponseDto
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Expect err to be nil. Instead got: %v\n", err)
				}
				if response.Pwd != tc.expectedMessage {
					t.Fatalf("Wrong password returned. Expected: %v got: %v\n", tc.expectedMessage, response.Pwd)
				}
			} else {
				var response dto.ErrorDto
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Expect err to be nil. Instead got: %v\n", err)
				}

				if response.Message != tc.expectedMessage {
					t.Fatalf("Wrong message returned. Expected: %v got: %v\n", tc.expectedMessage, response.Message)
				}
			}
		})
	}
}

func TestRemovePassword(t *testing.T) {
	cases := []struct {
		label          string
		key            string
		expectedStatus int
	}{
		{"TestRemovePassword", "orkut", http.StatusNoContent},
		{"TestRemovePasswordKeyNotFound", "github", http.StatusNotFound},
	}

	router := chi.NewRouter()
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdHandler := NewPasswordHandler(&PasswordServiceMock{})
			url := fmt.Sprintf("/password/%v", tc.key)

			request := httptest.NewRequest(http.MethodDelete, url, nil)
			request.Header.Add("userId", "1")
			rr := httptest.NewRecorder()
			router.Delete("/password/{key}", pwdHandler.RemoveByKey)
			router.ServeHTTP(rr, request)

			if rr.Code != tc.expectedStatus {
				t.Fatalf("Wrong status returned. Expected: %v got: %v\n", tc.expectedStatus, rr.Code)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	cases := []struct {
		label          string
		masterId       string
		expectedStatus int
	}{
		{"TestKeys", "1", http.StatusOK},
		{"TestKeysError", "", http.StatusInternalServerError},
	}

	router := chi.NewRouter()
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdHandler := NewPasswordHandler(&PasswordServiceMock{})

			request := httptest.NewRequest(http.MethodGet, "/password", nil)
			request.Header.Add("userId", tc.masterId)
			rr := httptest.NewRecorder()
			router.Get("/password", pwdHandler.Keys)
			router.ServeHTTP(rr, request)

			if rr.Code != tc.expectedStatus {
				t.Fatalf("Wrong status returned. Expected: %v got: %v\n", tc.expectedStatus, rr.Code)
			}
		})
	}
}

func TestUpdateKey(t *testing.T) {
	cases := []struct {
		label          string
		body           string
		expectedStatus int
	}{
		{"TestUpdateKey", `{"pwd": "123456", "key": "orkut"}`, http.StatusNoContent},
		{"TestUpdateKeyErrNotFound", `{"pwd": "123456", "key": "github"}`, http.StatusNotFound},
		{"TestUpdateKeyErrInvalidBody", `invalidbody`, http.StatusBadRequest},
	}

	router := chi.NewRouter()
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdHandler := NewPasswordHandler(&PasswordServiceMock{})

			request := httptest.NewRequest(http.MethodPut, "/password", strings.NewReader(tc.body))
			request.Header.Add("userId", "1")
			rr := httptest.NewRecorder()
			router.Put("/password", pwdHandler.UpdateByKey)
			router.ServeHTTP(rr, request)

			if rr.Code != tc.expectedStatus {
				t.Fatalf("Wrong status returned. Expected: %v got: %v\n", tc.expectedStatus, rr.Code)
			}
		})
	}
}
