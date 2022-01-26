package model

import (
	"database/sql"
)

type Master struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type MasterRepository interface {
	Save(*Master) error
	FindByEmail(string) (*Master, error)
}

type MasterRepositoryImpl struct {
	db *sql.DB
}

func NewMasterRepository(db *sql.DB) *MasterRepositoryImpl {
	return &MasterRepositoryImpl{db: db}
}

func (mr *MasterRepositoryImpl) Save() error {
	return nil
}

func (mr *MasterRepositoryImpl) FindByEmail(email string) (*Master, error) {
	return nil, nil
}
