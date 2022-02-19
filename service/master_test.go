package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/danilomarques1/gopmserver/dto"
	"github.com/danilomarques1/gopmserver/model"
	"github.com/danilomarques1/gopmserver/util"
)

const PWD_HASH = "$2y$04$PPdW8FkiC20NvMaXQcqCduB8ikmqD8Dm29t293/KMOl7iLQHI2/oa"

type masterRepositoryMock struct {
}

func (mrm *masterRepositoryMock) Save(master *model.Master) error {
	if len(master.Email) == 0 {
		return errors.New("Password should be not null")
	}
	return nil
}

func (mrm *masterRepositoryMock) FindByEmail(email string) (*model.Master, error) {
	if email == "test@mail.com" {
		return &model.Master{Id: "1", Email: "test@gmail.com", PwdHash: PWD_HASH}, nil
	}

	return nil, util.NewApiError("Invalid email", http.StatusNotFound)
}

func TestSaveMaster(t *testing.T) {
	cases := []struct {
		label     string
		masterDto *dto.MasterRequestDto
		err       bool // if err is returned
	}{
		{"TestSaveMaster", &dto.MasterRequestDto{Email: "fitz@mail.com", Pwd: "123456"}, false},
		{"TestSaveMasterErrEmailAlreadyInUse", &dto.MasterRequestDto{Email: "test@mail.com", Pwd: "12345"}, true},
		{"TestSaveMasterErrEmailAlreadyInUse", &dto.MasterRequestDto{Email: "", Pwd: "12345"}, true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			masterService := NewMasterService(&masterRepositoryMock{})
			err := masterService.Save(tc.masterDto)
			if (err != nil) != tc.err {
				t.Fatalf("Expect an error to be returned: %v\n", tc.err)
			}
		})
	}
}

func TestSessionMaster(t *testing.T) {
	cases := []struct {
		label      string
		sessionDto *dto.SessionRequestDto
		err        bool
	}{
		{"TestSessionMaster", &dto.SessionRequestDto{Email: "test@mail.com", Pwd: "123456"}, false},
		{"TestSessionMaster", &dto.SessionRequestDto{Email: "fitz@mail.com", Pwd: "123456"}, true},
	}
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			masterService := NewMasterService(&masterRepositoryMock{})
			_, err := masterService.Session(tc.sessionDto)
			if (err != nil) != tc.err {
				t.Fatalf("Expect an error to be returned: %v %v\n", tc.err, err)
			}
		})
	}
}
