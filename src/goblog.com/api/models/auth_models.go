package models

import (
	"database/sql"
	"errors"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authManager struct {
	db *sql.DB
}

type Authorizer interface {
	Authenticate(challenge LoginRequest) (bool, error)
}

func NewAuthorizer(db *sql.DB) (Authorizer, error) {
	if db == nil {
		return nil, errors.New("auth manager error: cannot accept nil database handle")
	}
	return authManager{db}, nil
}

func (am authManager) Authenticate(challenge LoginRequest) (bool, error) {
	return true, nil
}
