package model

type Master struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	PwdHash string `json:"pwd"`
}

type MasterRepository interface {
	// store a new master
	Save(*Master) error
	// find a master using the email
	FindByEmail(string) (*Master, error)
}
