package service

import (
	"log"
	"net/http"

	"github.com/danilomarques1/gopmserver/dto"
	"github.com/danilomarques1/gopmserver/model"
	"github.com/danilomarques1/gopmserver/util"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type MasterService struct {
	masterRepository model.MasterRepository
}

func NewMasterService(repository model.MasterRepository) *MasterService {
	return &MasterService{masterRepository: repository}
}

func (ms *MasterService) Save(masterDto *dto.MasterRequestDto) error {
	if _, err := ms.masterRepository.FindByEmail(masterDto.Email); err == nil {
		return util.NewApiError("E-mail already in use", http.StatusBadRequest)
	}
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(masterDto.Pwd), bcrypt.MinCost)
	if err != nil {
		return util.NewApiError(err.Error(), http.StatusInternalServerError)
	}

	master := model.Master{
		Id:      uuid.NewString(),
		Email:   masterDto.Email,
		PwdHash: string(pwdHash),
	}
	if err := ms.masterRepository.Save(&master); err != nil {
		log.Printf("Error saving repository %v\n", err)
		return util.NewApiError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (ms *MasterService) Session(sessionDto *dto.SessionRequestDto) (*dto.SessionResponseDto, error) {
	master, err := ms.masterRepository.FindByEmail(sessionDto.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(master.PwdHash), []byte(sessionDto.Pwd)); err != nil {
		return nil, util.NewApiError("Invalid password", http.StatusUnauthorized)
	}

	token, err := util.GenToken(master.Id)
	if err != nil {
		return nil, err
	}
	response := dto.SessionResponseDto{Token: token}
	return &response, nil
}
