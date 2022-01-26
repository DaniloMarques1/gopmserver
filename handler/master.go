package handler

import (
	"github.com/danilomarques1/gopmserver/service"
)

type MasterHandler struct {
	masterService *service.MasterService
}

func NewMasterHandler(masterService *service.MasterService) *MasterHandler {
	return &MasterHandler{masterService: masterService}
}
