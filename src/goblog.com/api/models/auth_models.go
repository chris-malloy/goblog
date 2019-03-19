package models

import (
	"database/sql"
	"errors"
)

type authManager struct {
	db *sql.DB
}

type Authorizer interface {
}

func NewAuthManager(db *sql.DB) (Authorizer, error) {
	if db == nil {
		return nil, errors.New("auth manager error: cannot accept nil database handle")
	}
	return authManager{db}, nil
}
