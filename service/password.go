package service

import (
	"net/http"

	"github.com/danilomarques1/gopmserver/dto"
	"github.com/danilomarques1/gopmserver/model"
	"github.com/danilomarques1/gopmserver/util"
	"github.com/google/uuid"
)

type PasswordService struct {
	pwdRepository model.PasswordInterface
}

func NewPasswordService(pwdRepository model.PasswordInterface) *PasswordService {
	return &PasswordService{pwdRepository: pwdRepository}
}

func (ps *PasswordService) Save(pwdDto dto.PasswordRequestDto) error {
	if _, err := ps.pwdRepository.FindByKey(pwdDto.Key); err == nil {
		return util.NewApiError("Key already in use", http.StatusBadRequest)
	}

	password := model.Password{
		Id:  uuid.NewString(),
		Key: pwdDto.Key,
		Pwd: pwdDto.Pwd,
	}
	if err := ps.pwdRepository.Save(&password); err != nil {
		return err
	}

	return nil
}
