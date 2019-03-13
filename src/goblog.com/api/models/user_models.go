package models

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// used to create new users as well as update them
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
	InsertUser(payload NewUserRequest) (*User, error)
	SelectUserById(userId int64) (*User, error)
	SelectUserByEmail(email string) (*User, error)
	UpdateUserById(userId int64, payload User) (*User, error)
	DeleteUserById(userId int64) (bool, error)
}

func NewUserManager(db *sql.DB) (UserCRUD, error) {
	if db == nil {
		return nil, errors.New("user manager error: cannot accept nil database handle")
	}
	return userManger{db}, nil
}

func (um userManger) InsertUser(newUser NewUserRequest) (*User, error) {
	query, err := um.db.Prepare(insertUserSQL)
	if err != nil {
		return nil, fmt.Errorf("insert user error: cannot prepare query: %v", err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword(bytes.NewBufferString(newUser.Password).Bytes(), 10)
	if err != nil {
		return nil, fmt.Errorf("insert user error: cannot generate password hash: %v", err.Error())
	}

	now := time.Now()
	results, err := query.Exec(newUser.FirstName, newUser.LastName, newUser.Email, hash, now.Format(time.RFC3339))
	if err != nil {
		return nil, fmt.Errorf("inster user error: cannot execute query: %v", err.Error())
	}

	if !areRowsAffected(results) {
		return nil, fmt.Errorf("insert user error: no rows affected while attempting to insert new user")
	} else {
		return um.SelectUserByEmail(newUser.Email)
	}
}

func (um userManger) SelectUserById(userId int64) (*User, error) {
	var who User
	err := um.db.QueryRow(selectUserByIdSQL, userId).Scan(
		&who.ID, &who.Email, &who.FirstName, &who.LastName, &who.LastSignedInAt, &who.SignInCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, fmt.Errorf("unexpected error while selecting user by id: %v", err.Error())
		}
	} else {
		return &who, nil
	}
}

func (um userManger) SelectUserByEmail(email string) (*User, error) {
	var who User
	err := um.db.QueryRow(selectUserByEmailSQL, email).Scan(
		&who.ID, &who.Email, &who.FirstName, &who.LastName, &who.LastSignedInAt, &who.SignInCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, fmt.Errorf("unexpected error while selecting user by email: %v", err.Error())
		}
	} else {
		return &who, nil
	}
}

func (um userManger) UpdateUserById(userId int64, payload User) (*User, error) {
	query, err := um.db.Prepare(updateUserByIdSQL)
	if err != nil {
		return nil, fmt.Errorf("update user by id error: can't prepare query: %v", err.Error())
	}

	results, err := query.Exec(userId, payload.Email, payload.FirstName, payload.LastName)
	if err != nil {
		return nil, fmt.Errorf("update user by id error: cannot execute query: %v", err.Error())
	}

	if !areRowsAffected(results) {
		return nil, fmt.Errorf("update user by id error: no rows affected while attempting to update user by id: user may not exist")
	} else {
		return um.SelectUserById(userId)
	}
}

func (um userManger) DeleteUserById(userId int64) (bool, error) {
	query, err := um.db.Prepare(deleteUserByIdSQL)
	if err != nil {
		return false, fmt.Errorf("delete user by id error: can't prepare query: %v", err.Error())
	}

	results, err := query.Exec(userId)
	if err != nil {
		return false, fmt.Errorf("delete user by id error: cannot execute query: %v", err.Error())
	}

	if !areRowsAffected(results) {
		return false, fmt.Errorf("delete user by id error: no rows affected while attempting to delete user by id: user may not exist")
	} else {
		return true, nil
	}
}

func areRowsAffected(results sql.Result) bool {
	count, _ := results.RowsAffected()
	return count == 1
}
