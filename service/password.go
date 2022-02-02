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
	return &dto.PasswordResponseDto{
		Id:  password.Id,
		Key: password.Key,
		Pwd: password.Pwd,
	}, nil
}

func (ps *PasswordService) Keys(masterId string) (*dto.PasswordKeysDto, error) {
	keys, err := ps.pwdRepository.Keys(masterId)
	if err != nil {
		return nil, err
	}
	return &dto.PasswordKeysDto{Keys: keys}, nil
}

func (ps *PasswordService) RemoveByKey(masterId, key string) error {
	password, err := ps.pwdRepository.FindByKey(masterId, key)
	if err != nil {
		return err
	}
	if err := ps.pwdRepository.RemoveByKey(masterId, password.Key); err != nil {
		return err
	}
	return nil
}

func (ps *PasswordService) UpdateByKey(masterId string, pwdDto *dto.PasswordUpdateRequestDto) error {
	password, err := ps.pwdRepository.FindByKey(masterId, pwdDto.Key)
	if err != nil {
		return err
	}
	password.Pwd = pwdDto.Pwd
	if err := ps.pwdRepository.UpdateByKey(masterId, password); err != nil {
		return err
	}
	return nil
}
