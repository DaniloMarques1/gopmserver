package dto

type PasswordRequestDto struct {
	Key string `json:"key"`
	Pwd string `json:"pwd"`
}

// TODO could be different
type PasswordResponseDto struct {
	Id  string `json:"id"`
	Key string `json:"key"`
	Pwd string `json:"pwd"`
}

type PasswordUpdateRequestDto struct {
	Key string `json:"key"`
	Pwd string `json:"pwd"`
}

type PasswordKeysDto struct {
	Keys []string `json:"keys"`
}
