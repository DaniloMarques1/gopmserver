package model

type Password struct {
	Id       string `json:"id"`
	Key      string `json:"key"`
	Pwd      string `json:"pwd"`
	MasterId string `json:"master_id"`
}

type PasswordInterface interface {
	Save(password *Password) error
	FindByKey(masterId string, key string) (*Password, error)
	FindAll(masterId string) ([]Password, error)
	Keys(masterId string) ([]string, error)
	RemoveByKey(masterId string, key string) error
}
