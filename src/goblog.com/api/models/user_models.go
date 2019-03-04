package models

import (
	"database/sql"
	"errors"
)

type NewUserRequest struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

type userManger struct {
	db *sql.DB
}

type Users interface {
}

func NewUserManager(db *sql.DB) (Users, error) {
	if db == nil {
		return nil, errors.New("cannot accept nil database handle")
	}
	return userManger{db}, nil
}
