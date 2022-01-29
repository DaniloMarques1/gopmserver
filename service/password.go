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

func (ps *PasswordService) Save(masterId string, pwdDto dto.PasswordRequestDto) error {
	if _, err := ps.pwdRepository.FindByKey(masterId, pwdDto.Key); err == nil {
		return util.NewApiError("Key already in use", http.StatusBadRequest)
	}

	password := model.Password{
		Id:       uuid.NewString(),
		Key:      pwdDto.Key,
		Pwd:      pwdDto.Pwd,
		MasterId: masterId,
	}
	if err := ps.pwdRepository.Save(&password); err != nil {
		return err
	}

	return nil
}

func (ps *PasswordService) FindByKey(masterId, key string) (*dto.PasswordResponseDto, error) {
	password, err := ps.pwdRepository.FindByKey(masterId, key)
	if err != nil {
		return nil, err
	}
	response := dto.PasswordResponseDto{
		Id:  password.Id,
		Key: password.Key,
		Pwd: password.Pwd,
	}

	return &response, nil
}
