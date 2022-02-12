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
	masterHandler := NewMasterHandler(&MasterServiceMock{})
	body := `{"email": "test@mail.com", "pwd": "secret"}`
	request, err := http.NewRequest(http.MethodPost, "/master", strings.NewReader(body))
	if err != nil {
		t.Fatalf("Error creating the request. Got: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(masterHandler.Save)
	handler.ServeHTTP(rr, request)
	if rr.Code != http.StatusCreated {
		t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", http.StatusCreated, rr.Code)
	}
}

func TestSaveMasterErrEmailAlreadyInUse(t *testing.T) {
	masterHandler := NewMasterHandler(&MasterServiceMock{})
	body := `{"email": "fitz@mail.com", "pwd": "secret"}`
	request, err := http.NewRequest(http.MethodPost, "/master", strings.NewReader(body))
	if err != nil {
		t.Fatalf("Error creating the request. Got: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(masterHandler.Save)
	handler.ServeHTTP(rr, request)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", http.StatusBadRequest, rr.Code)
	}
}

func TestSaveMasterErrInvalidBody(t *testing.T) {
	masterHandler := NewMasterHandler(&MasterServiceMock{})
	body := `thisisnotavalidjson`
	request, err := http.NewRequest(http.MethodPost, "/master", strings.NewReader(body))
	if err != nil {
		t.Fatalf("Error creating the request. Got: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(masterHandler.Save)
	handler.ServeHTTP(rr, request)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", http.StatusBadRequest, rr.Code)
	}
}

func TestSessionMaster(t *testing.T) {
	masterHandler := NewMasterHandler(&MasterServiceMock{})
	body := `{"email": "fitz@mail.com", "pwd": "123456"}`
	request, err := http.NewRequest(http.MethodPost, "/session", strings.NewReader(body))
	if err != nil {
		t.Fatalf("Expected error to be nil instead got: %v\n", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(masterHandler.Session)
	handler.ServeHTTP(rr, request)

	if rr.Code != http.StatusOK {
		t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", http.StatusOK, rr.Code)
	}

	var response dto.SessionResponseDto
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Expect err to be nil when parsing response. Instead got: %v\n", err)
	}

	if response.Token != "token" {
		t.Fatalf("Wrong token returned: %v\n", response.Token)
	}
}

func TestSessionMasterErrInvalidPwd(t *testing.T) {
	masterHandler := NewMasterHandler(&MasterServiceMock{})
	body := `{"email": "fitz@mail.com", "pwd": "wrongpassword"}`
	request, err := http.NewRequest(http.MethodPost, "/session", strings.NewReader(body))
	if err != nil {
		t.Fatalf("Expected error to be nil instead got: %v\n", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(masterHandler.Session)
	handler.ServeHTTP(rr, request)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", http.StatusUnauthorized, rr.Code)
	}

	var response dto.ErrorDto
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Expect err to be nil when parsing response. Instead got: %v\n", err)
	}

	if response.Message != "Invalid password" {
		t.Fatalf("Wrong message returned: %v\n", response.Message)
	}
}

func TestSessionMasterErrInvalidJson(t *testing.T ) {
	masterHandler := NewMasterHandler(&MasterServiceMock{})
	body := `invalidjson`
	request, err := http.NewRequest(http.MethodPost, "/session", strings.NewReader(body))
	if err != nil {
		t.Fatalf("Expected error to be nil instead got: %v\n", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(masterHandler.Session)
	handler.ServeHTTP(rr, request)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("Wrong status code returned. Expect: %v got: %v\n", http.StatusBadRequest, rr.Code)
	}

	var response dto.ErrorDto
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Expect err to be nil when parsing response. Instead got: %v\n", err)
	}

	if response.Message != "Invalid body" {
		t.Fatalf("Wrong message returned: %v\n", response.Message)
	}
}
