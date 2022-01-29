package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/danilomarques1/gopmserver/dto"
	"github.com/danilomarques1/gopmserver/service"
	"github.com/danilomarques1/gopmserver/util"
)

type MasterHandler struct {
	masterService *service.MasterService
}

func NewMasterHandler(masterService *service.MasterService) *MasterHandler {
	return &MasterHandler{masterService: masterService}
}

func (mh *MasterHandler) Save(w http.ResponseWriter, r *http.Request) {
	var masterDto dto.MasterRequestDto
	if err := json.NewDecoder(r.Body).Decode(&masterDto); err != nil {
		log.Printf("ERR parsing body: %v\n", err)
		util.RespondJSON(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := mh.masterService.Save(&masterDto); err != nil {
		log.Printf("Error saving %v\n", err)
		util.RespondERR(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (mh *MasterHandler) Session(w http.ResponseWriter, r *http.Request) {
	var sessionDto dto.SessionRequestDto
	if err := json.NewDecoder(r.Body).Decode(&sessionDto); err != nil {
		log.Printf("ERR parsing body: %v\n", err)
		util.RespondJSON(w, "Invalid body", http.StatusBadRequest)
		return
	}
	response, err := mh.masterService.Session(&sessionDto)
	if err != nil {
		log.Printf("Error %v\n", err)
		util.RespondERR(w, err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (mg *MasterHandler) GetPassword(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("userId")
	fmt.Println(id)

	w.Write([]byte("Hello\n"))
}
