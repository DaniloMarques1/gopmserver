package dto

type MasterRequestDto struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type SessionRequestDto struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type SessionResponseDto struct {
	Token string `json:"token"`
}
