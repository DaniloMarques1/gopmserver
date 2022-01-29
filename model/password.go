package model

type Password struct {
	Id  string `json:"id"`
	Key string `json:"key"`
	Pwd string `json:"pwd"`
}

type PasswordInterface interface {
	Save(*Password) error
	FindByKey(string) (*Password, error)
	FindAll() ([]Password, error)
	Keys() ([]string, error)
	RemoveByKey(string) error
}
