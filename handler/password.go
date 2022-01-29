package handler

import (
	"encoding/json"
	"net/http"

	"github.com/danilomarques1/gopmserver/dto"
	"github.com/danilomarques1/gopmserver/service"
	"github.com/danilomarques1/gopmserver/util"
)

type PasswordHandler struct {
	pwdService *service.PasswordService
}

func NewPasswordHandler(pwdService *service.PasswordService) *PasswordHandler {
	return &PasswordHandler{pwdService: pwdService}
}

func (ph *PasswordHandler) Save(w http.ResponseWriter, r *http.Request) {
	var pwdDto dto.PasswordRequestDto
	if err := json.NewDecoder(r.Body).Decode(&pwdDto); err != nil {
		util.RespondJSON(w, "Invalid json", http.StatusBadRequest)
		return
	}
	if err := ph.pwdService.Save(pwdDto); err != nil {
		util.RespondERR(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
