package db

import (
	"errors"
	"fmt"
	"os"
)

type Creds struct {
	host    string
	user    string
	pass    string
	dbname  string
	options string
}

func GetCredsFromEnv() (*Creds, error) {
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, errors.New("no host specified in environment")
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return nil, errors.New("no user specified in environment")
	}

	pass, ok := os.LookupEnv("DB_PASS")
	if !ok {
		return nil, errors.New("no password specified in environment")
	}

	dbname, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, errors.New("no database specified in environment")
	}

	options := os.Getenv("DB_OPTN")

	return &Creds{host, user, pass, dbname, options}, nil
}

const dbConnectionTemplate = `postgres://%s:%s@%s:5432/%s`

func (c Creds) ToConnectionString() string {
	connStr := fmt.Sprintf(dbConnectionTemplate, c.user, c.pass, c.host, c.dbname)

	connStr = appendOptions(connStr, c.options)

	return connStr
}

func appendOptions(connStr string, options string) string {
	if len(options) > 0 {
		connStr += "?" + options
	}
	return connStr
}
