package model

type Master struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	PwdHash string `json:"pwd"`
}

type MasterRepository interface {
	Save(*Master) error
	FindByEmail(string) (*Master, error)
}
