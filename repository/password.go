package repository

import (
	"database/sql"

	"github.com/danilomarques1/gopmserver/model"
)

type PasswordRepositoryImpl struct {
	db *sql.DB
}

func NewPasswordRepository(db *sql.DB) *PasswordRepositoryImpl {
	return &PasswordRepositoryImpl{db: db}
}

func (pr *PasswordRepositoryImpl) Save(password *model.Password) error {
	stmt, err := pr.db.Prepare("insert into password(id, key, pwd, master_id) values($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(password.Id, password.Key, password.Pwd, password.MasterId)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PasswordRepositoryImpl) FindByKey(masterId, key string) (*model.Password, error) {
	stmt, err := pr.db.Prepare("select id, key, pwd from password where key = $1 and master_id = $2")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var password model.Password
	err = stmt.QueryRow(key, masterId).Scan(&password.Id, &password.Key, &password.Pwd)
	if err != nil {
		return nil, err
	}

	return &password, nil
}

func (pr *PasswordRepositoryImpl) FindAll(masterId string) ([]model.Password, error) {
	stmt, err := pr.db.Prepare("select id, key, pwd from password")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return nil, nil
}

func (pr *PasswordRepositoryImpl) Keys(masterId string) ([]string, error) {
	return nil, nil
}

func (pr *PasswordRepositoryImpl) RemoveByKey(masterId, key string) error {
	stmt, err := pr.db.Prepare("delete from password where key = $1 and master_id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(key, masterId)
	if err != nil {
		return err
	}

	return nil
}
