package models

import "database/sql"

type Users interface {
}

func NewUserManager(db *sql.DB) (Users, error) {
	return db, nil
}
