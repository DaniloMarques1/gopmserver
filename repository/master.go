package repository

import (
	"database/sql"
	"github.com/danilomarques1/gopmserver/model"
)

type MasterRepositoryImpl struct {
	db *sql.DB
}

func NewMasterRepository(db *sql.DB) *MasterRepositoryImpl {
	return &MasterRepositoryImpl{db: db}
}

func (mr *MasterRepositoryImpl) Save(master *model.Master) error {
	return nil
}

func (mr *MasterRepositoryImpl) FindByEmail(email string) (*model.Master, error) {
	return nil, nil
}
