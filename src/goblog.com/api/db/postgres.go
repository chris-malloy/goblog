package db

import "errors"

type Creds struct {
}

func GetCredsFromEnv() (*Creds, error) {
	return &Creds{}, errors.New("no creds found")
}
