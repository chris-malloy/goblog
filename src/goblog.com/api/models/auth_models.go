package models

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	var encryptedPassword string
	err := am.db.QueryRow(getEncryptedPasswordSQL, challenge.Email).Scan(&encryptedPassword)
	if err != nil {
		return false, fmt.Errorf("authentication error: %v", err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(challenge.Password))
	if err != nil {
		return false, fmt.Errorf("authentication error: %v", err.Error())
	}

	return true, nil
}
