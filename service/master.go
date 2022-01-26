package service

import (
	"github.com/danilomarques1/gopmserver/model"
)

type MasterService struct {
	masterRepository model.MasterRepository
}

func NewMasterService(repository model.MasterRepository) *MasterService {
	return &MasterService{masterRepository: repository}
}

// return specific api error
func (ms *MasterService) Save(master *model.Master) error {
	return nil
}
