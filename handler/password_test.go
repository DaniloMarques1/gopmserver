package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danilomarques1/gopmserver/dto"
	"github.com/danilomarques1/gopmserver/util"
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
	return nil, nil
}

func (psm *PasswordServiceMock) Keys(masterId string) (*dto.PasswordKeysDto, error) {
	return nil, nil
}

func (psm *PasswordServiceMock) RemoveByKey(masterId, key string) error {
	return nil
}

func (psm *PasswordServiceMock) UpdateByKey(masterId string, pwdDto *dto.PasswordUpdateRequestDto) error {
	return nil
}

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

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdHandler := NewPasswordHandler(&PasswordServiceMock{})
			request, err := http.NewRequest(http.MethodPost, "/password", strings.NewReader(tc.body))
			if err != nil {
				t.Fatalf("Expect err to be nil. Instead got: %v\n", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(pwdHandler.Save)
			handler.ServeHTTP(rr, request)

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
