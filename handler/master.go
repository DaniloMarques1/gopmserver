package handler

import (
	"net/http"

	"github.com/danilomarques1/gopmserver/service"
)

type MasterHandler struct {
	masterService *service.MasterService
}

func NewMasterHandler(masterService *service.MasterService) *MasterHandler {
	return &MasterHandler{masterService: masterService}
}

func (mh *MasterHandler) Save(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}
