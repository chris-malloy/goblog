package models

import "database/sql"

type authManager struct {
	db *sql.DB
}

type Authorizer interface {
}

func NewAuthManager(db *sql.DB) (Authorizer, error) {
	return authManager{db}, nil
}
