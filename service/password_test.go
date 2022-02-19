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

var password = &model.Password{
	Id:       "1",
	Key:      "github",
	Pwd:      "123456",
	MasterId: "2",
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
	cases := []struct{
		label         string
		pwdDto        *dto.PasswordRequestDto
		masterId      string
		returnedError bool // if returns an error
	}{
		{"ShouldSavePassword", &dto.PasswordRequestDto{Key: "newKey", Pwd: "123456"}, "1", false},
		{"ShouldSavePasswordErroKeyInUse", &dto.PasswordRequestDto{Key: "github", Pwd: "123456"}, "1", true},
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
