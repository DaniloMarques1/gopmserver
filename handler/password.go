package handler

import (
	"encoding/json"
	"net/http"

	"github.com/danilomarques1/gopmserver/dto"
	"github.com/danilomarques1/gopmserver/service"
	"github.com/danilomarques1/gopmserver/util"

	"github.com/go-chi/chi/v5"
)

type PasswordHandler struct {
	pwdService service.PasswordService
}

func NewPasswordHandler(pwdService service.PasswordService) *PasswordHandler {
	return &PasswordHandler{pwdService: pwdService}
}

func (ph *PasswordHandler) Save(w http.ResponseWriter, r *http.Request) {
	var pwdDto dto.PasswordRequestDto
	if err := json.NewDecoder(r.Body).Decode(&pwdDto); err != nil {
		util.RespondJSON(w, "Invalid json", http.StatusBadRequest)
		return
	}
	masterId := r.Header.Get("userId")
	if err := ph.pwdService.Save(masterId, &pwdDto); err != nil {
		util.RespondERR(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ph *PasswordHandler) FindByKey(w http.ResponseWriter, r *http.Request) {
	masterId := r.Header.Get("userId")
	key := chi.URLParam(r, "key")
	response, err := ph.pwdService.FindByKey(masterId, key)
	if err != nil {
		util.RespondERR(w, err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (ph *PasswordHandler) RemoveByKey(w http.ResponseWriter, r *http.Request) {
	masterId := r.Header.Get("userId")
	key := chi.URLParam(r, "key")
	if err := ph.pwdService.RemoveByKey(masterId, key); err != nil {
		util.RespondERR(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ph *PasswordHandler) Keys(w http.ResponseWriter, r *http.Request) {
	masterId := r.Header.Get("userId")
	response, err := ph.pwdService.Keys(masterId)
	if err != nil {
		util.RespondERR(w, err)
		return
	}
	json.NewEncoder(w).Encode(response)
}
