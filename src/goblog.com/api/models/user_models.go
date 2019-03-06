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
	InsertUser(newUser NewUserRequest) (*User, error)
	SelectUserByEmail(email string) (*User, error)
}

func NewUserManager(db *sql.DB) (UserCRUD, error) {
	if db == nil {
		return nil, errors.New("cannot accept nil database handle")
	}
	return userManger{db}, nil
}

func (um userManger) InsertUser(newUser NewUserRequest) (*User, error) {
	query, err := um.db.Prepare(insertUserSQL)
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
		return um.SelectUserByEmail(newUser.Email)
	}
}

func (um userManger) SelectUserByEmail(email string) (*User, error) {
	var who User
	err := um.db.QueryRow(selectUserSQL, email).Scan(
		&who.ID, &who.Email, &who.FirstName, &who.LastName, &who.LastSignedInAt, &who.SignInCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	} else {
		return &who, nil
	}
}
