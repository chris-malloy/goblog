package models

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID             int64      `json:"id"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Email          string     `json:"email"`
	LastSignedInAt *time.Time `json:"last_sign_in"`
	SignInCount    int        `json:"sign_in_count"`
}

// This has to look different because we don't ship a password
// with every user. So, a new user request will be shaped differently
// than a User request.
type NewUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type userManger struct {
	db *sql.DB
}

type UserCRUD interface {
	CreateUser(newUser NewUserRequest) (*User, error)
}

func NewUserManager(db *sql.DB) (UserCRUD, error) {
	if db == nil {
		return nil, errors.New("cannot accept nil database handle")
	}
	return userManger{db}, nil
}

const createUserSQL = `
    INSERT INTO users (first_name, last_name, email, encrypted_password, sign_in_count, last_sign_in_at, created_time_stamp, updated_time_stamp)
    VALUES ($1, $2, $3, $4, 0, null, $5, null);
`

func (um userManger) CreateUser(newUser NewUserRequest) (*User, error) {
	query, err := um.db.Prepare(createUserSQL)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword(bytes.NewBufferString(newUser.Password).Bytes(), 10)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	results, err := query.Exec(newUser.FirstName, newUser.LastName, newUser.Email, hash, now.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	if count, err := results.RowsAffected(); count != 1 {
		return nil, fmt.Errorf("error: %v while creating user. No rows affected", err.Error())
	} else {
		return &User{FirstName: newUser.FirstName, LastName: newUser.LastName, Email: newUser.Email}, nil
	}
}
