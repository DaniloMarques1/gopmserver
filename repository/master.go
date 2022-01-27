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
	stmt, err := mr.db.Prepare("insert into master(id, email, pwd_hash) values($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(master.Id, master.Email, master.PwdHash)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MasterRepositoryImpl) FindByEmail(email string) (*model.Master, error) {
	stmt, err := mr.db.Prepare("select id, email, pwd_hash from master where email = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var master model.Master
	if err := stmt.QueryRow(email).Scan(&master.Id, &master.Email, &master.PwdHash); err != nil {
		return nil, err
	}
	return &master, nil
}
