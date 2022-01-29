package dto

type PasswordRequestDto struct {
	Key string `json:"key"`
	Pwd string `json:"pwd"`
}

type PasswordResponseDto struct {
	Id  string `json:"id"`
	Key string `json:"key"`
	Pwd string `json:"pwd"`
}

type PasswordKeysDto struct {
	Keys []string `json:"keys"`
}
