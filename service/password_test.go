package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/danilomarques1/gopmserver/dto"
	"github.com/danilomarques1/gopmserver/model"
	"github.com/danilomarques1/gopmserver/util"
)

type passwordRepositoryMock struct {
}

const (
	PWD_ID = "1"
	MASTER_ID = "2"
)

var password = &model.Password{
	Id:       PWD_ID,
	Key:      "github",
	Pwd:      "123456",
	MasterId: MASTER_ID,
}

func (prm *passwordRepositoryMock) Save(password *model.Password) error {
	if password.Key == "github" {
		return util.NewApiError("Key already in use", http.StatusBadRequest)
	}
	return nil
}

func (prm *passwordRepositoryMock) FindByKey(masterId, key string) (*model.Password, error) {
	if key != "github" {
		return nil, util.NewApiError("No password found with the given key", http.StatusNotFound)
	}
	return password, nil
}

func (prm *passwordRepositoryMock) FindAll(masterId string) ([]model.Password, error) {
	if len(masterId) == 0 {
		return nil, util.NewApiError("No password found with the given key", http.StatusNotFound)
	}
	passwords := []model.Password{*password}
	return passwords, nil
}

func (prm *passwordRepositoryMock) Keys(masterId string) ([]string, error) {
	// TODO
	if masterId == "3" {
		return nil, errors.New("Some error")
	}
	keys := []string{"github"}
	return keys, nil
}

func (prm *passwordRepositoryMock) RemoveByKey(masterId, key string) error {
	if len(key) == 0 {
		return errors.New("Some error")
	}
	return nil
}

func (prm *passwordRepositoryMock) UpdateByKey(masterId string, password *model.Password) error {
	// TODO
	if len(masterId) == 0 {
		return errors.New("Some error")
	}
	return nil
}

func TestSavePassword(t *testing.T) {
	cases := []struct {
		label         string
		pwdDto        *dto.PasswordRequestDto
		masterId      string
		returnedError bool // if returns an error
	}{
		{"ShouldSavePassword", &dto.PasswordRequestDto{Key: "newKey", Pwd: "123456"}, MASTER_ID, false},
		{"ShouldSavePasswordErroKeyInUse", &dto.PasswordRequestDto{Key: "github", Pwd: "123456"}, MASTER_ID, true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdService := NewPasswordService(&passwordRepositoryMock{})
			err := pwdService.Save(tc.masterId, tc.pwdDto)
			returnedError := err != nil
			if returnedError != tc.returnedError {
				t.Fatalf("Should return an error: %v instead got %v\n", tc.returnedError, err)
			}
		})
	}
}

func TestFindByKey(t *testing.T) {
	cases := []struct {
		label         string
		key           string
		masterId      string
		response      *dto.PasswordResponseDto
		returnedError bool // if returns an error
	}{
		{"ShouldFindTheKey", "github", MASTER_ID, &dto.PasswordResponseDto{Id: password.Id, Key: password.Key, Pwd: password.Pwd}, false},
		{"ShouldNotFindTheKey", "wrongkey", MASTER_ID, nil, true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdService := NewPasswordService(&passwordRepositoryMock{})
			response, err := pwdService.FindByKey(tc.masterId, tc.key)
			returnedErr := err != nil
			if returnedErr != tc.returnedError {
				t.Fatalf("Should return error: %v. Instead got: %v\n", tc.returnedError, err)
			}
			
			if tc.response != nil {
				if response == nil {
					t.Fatal("Response should not be nil\n")
				}
			} else {
				if response != nil {
					t.Fatalf("Response should be nil, instead got: %v\n", response)
				}
			}
		})
	}
}

func TestKeys(t *testing.T) {
	cases := []struct {
		label         string
		masterId      string
		response      *dto.PasswordKeysDto
		returnedError bool // if returns an error
	}{
		{"ShouldReturnKeys", MASTER_ID, &dto.PasswordKeysDto{Keys: []string{"github"}}, false},
		{"ShouldReturnError", "3", nil, true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdService := NewPasswordService(&passwordRepositoryMock{})
			response, err := pwdService.Keys(tc.masterId)
			returnedErr := err != nil
			if returnedErr != tc.returnedError {
				t.Fatalf("Should return error: %v. Instead got: %v\n", tc.returnedError, err)
			}
			
			if tc.response != nil {
				if response == nil {
					t.Fatal("Response should not be nil\n")
				}
			} else {
				if response != nil {
					t.Fatalf("Response should be nil, instead got: %v\n", response)
				}
			}
		})
	}
}

func TestRemoveByKey(t *testing.T) {
	cases := []struct {
		label         string
		masterId      string
		key           string
		returnedError bool // if returns an error
	}{
		{"ShouldRemoveKey", MASTER_ID, "github", false},
		{"ShouldNotRemoveKey", MASTER_ID, "", true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdService := NewPasswordService(&passwordRepositoryMock{})
			err := pwdService.RemoveByKey(tc.masterId, tc.key)
			returnedErr := err != nil
			if returnedErr != tc.returnedError {
				t.Fatalf("Should return error: %v. Instead got: %v\n", tc.returnedError, err)
			}
		})
	}
}

func TestUpdateByKey(t *testing.T) {
	cases := []struct {
		label         string
		masterId      string
		updateDto     *dto.PasswordUpdateRequestDto
		returnedError bool // if returns an error
	}{
		{"ShouldRemoveKey", MASTER_ID, &dto.PasswordUpdateRequestDto{Key: "github", Pwd: "oldOne"}, false},
		{"ShouldNotRemoveKey", "", &dto.PasswordUpdateRequestDto{Key: "", Pwd: "oldOne"}, true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdService := NewPasswordService(&passwordRepositoryMock{})
			err := pwdService.UpdateByKey(tc.masterId, tc.updateDto)
			returnedErr := err != nil
			if returnedErr != tc.returnedError {
				t.Fatalf("Should return error: %v. Instead got: %v\n", tc.returnedError, err)
			}
		})
	}
}
